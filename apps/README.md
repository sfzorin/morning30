# morning30 apps for iPhone and Android

Both apps show morning30.com full screen and add what only a phone app can do.
New programs, translations and fixes reach the apps as soon as the site is
deployed; a store update is needed only when the native part changes.

Shared with the site:

- `assets/app-bridge.js` — tells the site it runs inside an app, shows the
  app-only "Reminder" row in Settings and hands the chosen time to the phone.
- Settings → Delete account — required by both stores; works everywhere.
- `/privacy` — the privacy policy for the store listings.
- `/.well-known/assetlinks.json` — set `ANDROID_CERT_SHA256` on the server.
- Finished workouts: `assets/player.js` fires `m30:workout` (start, end, program
  title, kind strength/stretch); `app-bridge.js` hands it to the phone once
  (workouts shorter than a minute are skipped), which writes it to Apple Health
  or Health Connect. The apps only write workouts and never read health data.

## Android (`apps/android`)

A Trusted Web Activity (Chrome full screen) plus a small native part:

- `LauncherActivity` opens the site with `?app=android&reminder=HH:MM`;
- `ReminderActivity` receives the time from the site
  (`intent://reminder?…#Intent;scheme=morning30;…`), asks for notification
  permission and schedules the reminder;
- `ReminderReceiver` shows the notification and restores it after a reboot;
- `HealthActivity` receives a finished workout
  (`intent://health?start=…&end=…&kind=…&title=…`) and writes an
  `ExerciseSessionRecord` (calisthenics or stretching) to Health Connect,
  asking for the write-exercise permission the first time. Chrome opens
  `intent://` links only right after a tap, so a workout that ends on a timer
  is sent on the next tap (the done screen's button);
- `PrivacyActivity` is what Health Connect opens from "privacy policy"
  (shows https://morning30.com/privacy).

Health Connect needs `minSdk 26`, `compileSdk 36` and Android Gradle plugin
8.10+. In Play Console fill in the Health Connect declaration: the app writes
Exercise only, for workout history in the user's health app.

Build (JDK 17, Android SDK):

```bash
cd apps/android
JAVA_HOME=/opt/homebrew/opt/openjdk@17/libexec/openjdk.jdk/Contents/Home ./gradlew assembleDebug   # opens http://10.0.2.2:8080 (dev server)
JAVA_HOME=/opt/homebrew/opt/openjdk@17/libexec/openjdk.jdk/Contents/Home ./gradlew bundleRelease   # opens https://morning30.com
```

Release signing: create an upload key once and keep it outside the repository,

```bash
keytool -genkeypair -v -keystore ~/morning30-upload.jks -alias upload -keyalg RSA -keysize 2048 -validity 10000
```

then write `apps/android/keystore.properties` (ignored by git):

```
storeFile=/Users/you/morning30-upload.jks
storePassword=…
keyAlias=upload
keyPassword=…
```

Upload `app/build/outputs/bundle/release/app-release.aab` to Google Play
Console. In Play Console → Setup → App signing, copy the SHA-256 fingerprints
(the app signing key and the upload key) into `ANDROID_CERT_SHA256` on the
server, comma-separated; without it Android shows an address bar.

## iPhone (`apps/ios`)

A UIKit app around a WKWebView (`WebViewController`):

- injects `window.morning30App = {platform: "ios", reminder}` and receives the
  reminder through `webkit.messageHandlers.morning30`;
- schedules a daily local notification (`Reminders`);
- keeps the screen on during a workout (`/day/N`, `/start`);
- shows a "No internet connection" screen with a retry button;
- opens links to other sites in Safari; plays voice cues with the silent switch on;
- writes each finished workout to Apple Health (`Health.swift`, HealthKit
  entitlement in `project.yml`, usage texts in `Info.plist` and
  `*.lproj/InfoPlist.strings`); the permission sheet asks only to write
  Workouts. In App Store Connect add the HealthKit capability to the App ID
  and answer the privacy questions: Health & Fitness data, used for app
  functionality, written to Apple Health only.

Debug builds take `MORNING30_URL` (another dev server) and
`MORNING30_TEST_WORKOUT=1` (save a 15-minute test workout at launch), e.g.
`SIMCTL_CHILD_MORNING30_URL=http://localhost:8081/ xcrun simctl launch booted com.morning30.app`.

The Xcode project is generated from `project.yml`:

```bash
cd apps/ios && xcodegen generate
```

Debug builds open `http://localhost:8080`, release builds `https://morning30.com`.
To publish: open `Morning30.xcodeproj`, choose your team in Signing &
Capabilities, Product → Archive, then Distribute App → App Store Connect.

## Store listings

- Privacy policy URL: https://morning30.com/privacy
- Data collection: email and name (account), workout history and settings,
  and if the person fills them in: age, height, weight, concerns and fitness
  test results (app functionality); not linked to tracking; no third-party
  sharing. Workouts are written to Apple Health / Health Connect with
  permission; nothing is read from them.
- Account deletion: Settings → Delete account (in the app and on the site).
