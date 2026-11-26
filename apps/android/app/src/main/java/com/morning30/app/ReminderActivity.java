package com.morning30.app;

import android.Manifest;
import android.app.Activity;
import android.content.pm.PackageManager;
import android.net.Uri;
import android.os.Build;
import android.os.Bundle;

/**
 * Receives the reminder chosen on the site (morning30://reminder?time=07:30&title=…&body=…),
 * asks for notification permission when needed, schedules it and closes at once.
 */
public class ReminderActivity extends Activity {
    @Override
    protected void onCreate(Bundle state) {
        super.onCreate(state);
        Uri data = getIntent().getData();
        if (data == null) {
            finish();
            return;
        }
        String time = data.getQueryParameter("time");
        Reminders.save(this, time, data.getQueryParameter("title"), data.getQueryParameter("body"));
        Reminders.schedule(this);
        if (Reminders.parse(time) != null && Build.VERSION.SDK_INT >= 33
                && checkSelfPermission(Manifest.permission.POST_NOTIFICATIONS) != PackageManager.PERMISSION_GRANTED) {
            requestPermissions(new String[] {Manifest.permission.POST_NOTIFICATIONS}, 1);
            return; // closes in onRequestPermissionsResult
        }
        finish();
    }

    @Override
    public void onRequestPermissionsResult(int code, String[] permissions, int[] results) {
        finish();
    }
}
