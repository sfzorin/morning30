package com.morning30.app;

import android.net.Uri;

/** Opens morning30.com full screen and tells the site it runs inside the Android app. */
public class LauncherActivity extends com.google.androidbrowserhelper.trusted.LauncherActivity {
    @Override
    protected Uri getLaunchingUrl() {
        Uri url = super.getLaunchingUrl();
        return url.buildUpon()
                .appendQueryParameter("app", "android")
                .appendQueryParameter("reminder", Reminders.time(this))
                .build();
    }
}
