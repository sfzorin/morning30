import Foundation
import UserNotifications

/// The morning reminder: a daily local notification at the time chosen on the
/// site's settings page. It lives only on the phone; nothing goes to the server.
enum Reminders {
    private static let key = "reminder.time"
    private static let id = "morning"

    /// "HH:MM", or "" when the reminder is off.
    static var time: String { UserDefaults.standard.string(forKey: key) ?? "" }

    static func set(time: String, title: String, body: String) {
        UserDefaults.standard.set(time, forKey: key)
        let center = UNUserNotificationCenter.current()
        center.removePendingNotificationRequests(withIdentifiers: [id])
        let parts = time.split(separator: ":").compactMap { Int($0) }
        guard parts.count == 2, (0..<24).contains(parts[0]), (0..<60).contains(parts[1]) else { return }
        center.requestAuthorization(options: [.alert, .sound]) { granted, _ in
            guard granted else { return }
            let content = UNMutableNotificationContent()
            content.title = title.isEmpty ? "morning30" : title
            content.body = body
            content.sound = .default
            var when = DateComponents()
            when.hour = parts[0]
            when.minute = parts[1]
            let trigger = UNCalendarNotificationTrigger(dateMatching: when, repeats: true)
            center.add(UNNotificationRequest(identifier: id, content: content, trigger: trigger))
        }
    }
}
