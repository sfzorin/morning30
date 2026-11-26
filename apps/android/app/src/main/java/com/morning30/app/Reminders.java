package com.morning30.app;

import android.app.AlarmManager;
import android.app.NotificationChannel;
import android.app.NotificationManager;
import android.app.PendingIntent;
import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Build;

import java.util.Calendar;

/** The morning reminder: one daily notification at a time the user picked on the site. */
final class Reminders {
    static final String ACTION = "com.morning30.app.REMIND";
    private static final String PREFS = "reminder";
    private static final String CHANNEL = "reminder";

    private Reminders() {}

    static SharedPreferences prefs(Context c) {
        return c.getSharedPreferences(PREFS, Context.MODE_PRIVATE);
    }

    /** The saved time, "HH:MM", or "" when the reminder is off. */
    static String time(Context c) {
        return prefs(c).getString("time", "");
    }

    static void save(Context c, String time, String title, String body) {
        SharedPreferences.Editor e = prefs(c).edit().putString("time", time == null ? "" : time);
        if (title != null && !title.isEmpty()) e.putString("title", title);
        if (body != null && !body.isEmpty()) e.putString("body", body);
        e.apply();
    }

    private static PendingIntent alarmIntent(Context c) {
        Intent i = new Intent(c, ReminderReceiver.class).setAction(ACTION);
        return PendingIntent.getBroadcast(c, 0, i, PendingIntent.FLAG_UPDATE_CURRENT | PendingIntent.FLAG_IMMUTABLE);
    }

    /** Schedules the next reminder, or cancels it when the reminder is off. */
    static void schedule(Context c) {
        AlarmManager am = (AlarmManager) c.getSystemService(Context.ALARM_SERVICE);
        am.cancel(alarmIntent(c));
        String t = time(c);
        int[] hm = parse(t);
        if (hm == null) return;
        Calendar next = Calendar.getInstance();
        next.set(Calendar.HOUR_OF_DAY, hm[0]);
        next.set(Calendar.MINUTE, hm[1]);
        next.set(Calendar.SECOND, 0);
        next.set(Calendar.MILLISECOND, 0);
        if (!next.after(Calendar.getInstance())) next.add(Calendar.DAY_OF_YEAR, 1);
        // Inexact on purpose: no special alarm permission, and a few minutes either way is fine.
        am.setAndAllowWhileIdle(AlarmManager.RTC_WAKEUP, next.getTimeInMillis(), alarmIntent(c));
    }

    static int[] parse(String t) {
        if (t == null || !t.matches("\\d{2}:\\d{2}")) return null;
        int h = Integer.parseInt(t.substring(0, 2)), m = Integer.parseInt(t.substring(3));
        return h < 24 && m < 60 ? new int[] {h, m} : null;
    }

    static void ensureChannel(Context c) {
        if (Build.VERSION.SDK_INT < 26) return;
        NotificationManager nm = c.getSystemService(NotificationManager.class);
        if (nm.getNotificationChannel(CHANNEL) == null) {
            nm.createNotificationChannel(new NotificationChannel(CHANNEL, c.getString(R.string.channel_name), NotificationManager.IMPORTANCE_DEFAULT));
        }
    }

    static String channel() {
        return CHANNEL;
    }
}
