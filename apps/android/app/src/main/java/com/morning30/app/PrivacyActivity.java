package com.morning30.app;

import android.app.Activity;
import android.content.Intent;
import android.net.Uri;
import android.os.Bundle;

/**
 * Health Connect opens this screen when the person asks why the app wants
 * access; it shows the privacy policy on morning30.com.
 */
public class PrivacyActivity extends Activity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        try {
            startActivity(new Intent(Intent.ACTION_VIEW, Uri.parse("https://morning30.com/privacy")));
        } catch (Exception ignored) {
            // No browser: nothing else to show.
        }
        finish();
    }
}
