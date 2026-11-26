package com.morning30.app;

import android.Manifest;
import android.app.Notification;
import android.app.NotificationManager;
import android.app.PendingIntent;
import android.content.BroadcastReceiver;
import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.content.pm.PackageManager;
import android.os.Build;

/** Shows the morning notification and schedules the next one; restores it after a reboot. */
public class ReminderReceiver extends BroadcastReceiver {
    @Override
    public void onReceive(Context c, Intent intent) {
        if (Reminders.ACTION.equals(intent.getAction())) {
            show(c);
        }
        Reminders.schedule(c);
    }

    private static void show(Context c) {
        if (Build.VERSION.SDK_INT >= 33
                && c.checkSelfPermission(Manifest.permission.POST_NOTIFICATIONS) != PackageManager.PERMISSION_GRANTED) {
            return;
        }
        Reminders.ensureChannel(c);
        SharedPreferences p = Reminders.prefs(c);
        Intent open = new Intent(c, LauncherActivity.class).addFlags(Intent.FLAG_ACTIVITY_NEW_TASK);
        PendingIntent tap = PendingIntent.getActivity(c, 0, open, PendingIntent.FLAG_UPDATE_CURRENT | PendingIntent.FLAG_IMMUTABLE);
        Notification.Builder b = Build.VERSION.SDK_INT >= 26
                ? new Notification.Builder(c, Reminders.channel())
                : new Notification.Builder(c);
        Notification n = b.setSmallIcon(R.drawable.ic_notification)
                .setContentTitle(p.getString("title", c.getString(R.string.app_name)))
                .setContentText(p.getString("body", ""))
                .setContentIntent(tap)
                .setAutoCancel(true)
                .build();
        c.getSystemService(NotificationManager.class).notify(1, n);
    }
}
