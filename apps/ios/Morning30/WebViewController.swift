import UIKit
import WebKit

/// The whole app is morning30.com in a web view, plus what only a phone app
/// can do: the morning reminder, keeping the screen on during a workout, and a
/// clear "no internet" screen.
final class WebViewController: UIViewController, WKNavigationDelegate, WKUIDelegate, WKScriptMessageHandler {
    #if DEBUG
    // MORNING30_URL points a debug build at another dev server (simctl: SIMCTL_CHILD_MORNING30_URL).
    static let startURL = URL(string: ProcessInfo.processInfo.environment["MORNING30_URL"] ?? "http://localhost:8080/")!
    #else
    static let startURL = URL(string: "https://morning30.com/")!
    #endif

    private static let background = UIColor(red: 0x0F / 255, green: 0x17 / 255, blue: 0x2A / 255, alpha: 1)
    private var webView: WKWebView!
    private var urlObservation: NSKeyValueObservation?
    private let offline = OfflineView()

    override var preferredStatusBarStyle: UIStatusBarStyle { .lightContent }

    override func loadView() {
        let config = WKWebViewConfiguration()
        config.allowsInlineMediaPlayback = true
        config.mediaTypesRequiringUserActionForPlayback = []
        config.applicationNameForUserAgent = "morning30-ios"
        let scripts = WKUserContentController()
        // Tells the site it runs in the iPhone app (assets/app-bridge.js).
        let state = "window.morning30App = {platform: 'ios', reminder: '\(Reminders.time)'};"
        scripts.addUserScript(WKUserScript(source: state, injectionTime: .atDocumentStart, forMainFrameOnly: true))
        scripts.add(WeakHandler(self), name: "morning30")
        config.userContentController = scripts

        webView = WKWebView(frame: .zero, configuration: config)
        webView.navigationDelegate = self
        webView.uiDelegate = self
        webView.isOpaque = false
        webView.backgroundColor = Self.background
        webView.scrollView.backgroundColor = Self.background
        webView.allowsBackForwardNavigationGestures = true
        view = webView
    }

    override func viewDidLoad() {
        super.viewDidLoad()
        offline.isHidden = true
        offline.onRetry = { [weak self] in self?.reload() }
        view.addSubview(offline)
        offline.translatesAutoresizingMaskIntoConstraints = false
        NSLayoutConstraint.activate([
            offline.leadingAnchor.constraint(equalTo: view.leadingAnchor),
            offline.trailingAnchor.constraint(equalTo: view.trailingAnchor),
            offline.topAnchor.constraint(equalTo: view.topAnchor),
            offline.bottomAnchor.constraint(equalTo: view.bottomAnchor),
        ])
        // Keep the screen on while a workout plays (/day/N, /start) and during
        // the fitness self-test (/test).
        urlObservation = webView.observe(\.url, options: [.new]) { webView, _ in
            let path = webView.url?.path ?? ""
            UIApplication.shared.isIdleTimerDisabled = path.hasPrefix("/day/") || path == "/start" || path == "/test"
        }
        webView.load(URLRequest(url: Self.startURL))
        #if DEBUG
        // MORNING30_TEST_WORKOUT=1 saves a 15-minute test workout to Apple Health.
        if ProcessInfo.processInfo.environment["MORNING30_TEST_WORKOUT"] == "1" {
            Health.saveWorkout(start: Date().addingTimeInterval(-900), end: Date(), title: "morning30 test", kind: "strength")
        }
        #endif
    }

    private func reload() {
        offline.isHidden = true
        if webView.url == nil {
            webView.load(URLRequest(url: Self.startURL))
        } else {
            webView.reload()
        }
    }

    // MARK: Messages from the site

    func userContentController(_ controller: WKUserContentController, didReceive message: WKScriptMessage) {
        guard let body = message.body as? [String: Any] else { return }
        switch body["type"] as? String {
        case "reminder":
            Reminders.set(time: body["time"] as? String ?? "", title: body["title"] as? String ?? "", body: body["body"] as? String ?? "")
        case "workout":
            // A finished workout from assets/app-bridge.js: times are Unix milliseconds.
            guard let start = body["start"] as? Double, let end = body["end"] as? Double else { return }
            Health.saveWorkout(start: Date(timeIntervalSince1970: start / 1000), end: Date(timeIntervalSince1970: end / 1000),
                               title: body["title"] as? String ?? "", kind: body["kind"] as? String ?? "strength")
        default:
            break
        }
    }

    // MARK: Navigation

    func webView(_ webView: WKWebView, decidePolicyFor action: WKNavigationAction) async -> WKNavigationActionPolicy {
        guard let url = action.request.url else { return .allow }
        if url.host == Self.startURL.host || url.scheme == "about" || url.scheme == "blob" || url.scheme == "data" {
            return .allow
        }
        // Everything else (GitHub, mail links) opens outside the app.
        await UIApplication.shared.open(url)
        return .cancel
    }

    func webView(_ webView: WKWebView, createWebViewWith configuration: WKWebViewConfiguration, for action: WKNavigationAction, windowFeatures: WKWindowFeatures) -> WKWebView? {
        if let url = action.request.url {
            if url.host == Self.startURL.host {
                webView.load(action.request)
            } else {
                UIApplication.shared.open(url)
            }
        }
        return nil
    }

    func webView(_ webView: WKWebView, didFailProvisionalNavigation navigation: WKNavigation!, withError error: Error) {
        let code = (error as NSError).code
        if [NSURLErrorNotConnectedToInternet, NSURLErrorNetworkConnectionLost, NSURLErrorTimedOut,
            NSURLErrorCannotFindHost, NSURLErrorCannotConnectToHost].contains(code) {
            offline.isHidden = false
        }
    }

    func webView(_ webView: WKWebView, didFinish navigation: WKNavigation!) {
        offline.isHidden = true
    }

    func webViewWebContentProcessDidTerminate(_ webView: WKWebView) {
        webView.reload()
    }
}

/// Breaks the retain cycle between the web view's script handlers and the controller.
private final class WeakHandler: NSObject, WKScriptMessageHandler {
    weak var target: WKScriptMessageHandler?
    init(_ target: WKScriptMessageHandler) { self.target = target }
    func userContentController(_ controller: WKUserContentController, didReceive message: WKScriptMessage) {
        target?.userContentController(controller, didReceive: message)
    }
}

/// Shown instead of a blank page when the phone is offline.
private final class OfflineView: UIView {
    var onRetry: (() -> Void)?

    override init(frame: CGRect) {
        super.init(frame: frame)
        backgroundColor = UIColor(red: 0x0F / 255, green: 0x17 / 255, blue: 0x2A / 255, alpha: 1)
        let text = Self.texts()
        let label = UILabel()
        label.text = text.message
        label.textColor = .white
        label.font = .preferredFont(forTextStyle: .title3)
        label.numberOfLines = 0
        label.textAlignment = .center
        var config = UIButton.Configuration.filled()
        config.title = text.retry
        config.baseBackgroundColor = UIColor(red: 0xFB / 255, green: 0x92 / 255, blue: 0x3C / 255, alpha: 1)
        config.baseForegroundColor = .black
        config.cornerStyle = .capsule
        let button = UIButton(configuration: config, primaryAction: UIAction { [weak self] _ in self?.onRetry?() })
        let stack = UIStackView(arrangedSubviews: [label, button])
        stack.axis = .vertical
        stack.spacing = 20
        stack.alignment = .center
        addSubview(stack)
        stack.translatesAutoresizingMaskIntoConstraints = false
        NSLayoutConstraint.activate([
            stack.centerYAnchor.constraint(equalTo: centerYAnchor),
            stack.leadingAnchor.constraint(equalTo: layoutMarginsGuide.leadingAnchor),
            stack.trailingAnchor.constraint(equalTo: layoutMarginsGuide.trailingAnchor),
        ])
    }

    required init?(coder: NSCoder) { fatalError("not used") }

    private static func texts() -> (message: String, retry: String) {
        let lang = Locale.preferredLanguages.first.map { String($0.prefix(2)) } ?? "en"
        let all: [String: (String, String)] = [
            "en": ("No internet connection", "Try again"),
            "ru": ("Нет подключения к интернету", "Повторить"),
            "de": ("Keine Internetverbindung", "Erneut versuchen"),
            "es": ("Sin conexión a internet", "Reintentar"),
            "fr": ("Pas de connexion Internet", "Réessayer"),
            "it": ("Nessuna connessione a Internet", "Riprova"),
            "pt": ("Sem conexão com a internet", "Tentar de novo"),
            "tr": ("İnternet bağlantısı yok", "Tekrar dene"),
            "pl": ("Brak połączenia z internetem", "Spróbuj ponownie"),
            "ja": ("インターネットに接続されていません", "再試行"),
            "ko": ("인터넷에 연결되어 있지 않습니다", "다시 시도"),
            "zh": ("没有网络连接", "重试"),
            "ar": ("لا يوجد اتصال بالإنترنت", "إعادة المحاولة"),
        ]
        return all[lang] ?? all["en"]!
    }
}
