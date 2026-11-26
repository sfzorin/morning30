package com.morning30.app

import android.net.Uri
import android.os.Bundle
import android.util.Log
import androidx.activity.ComponentActivity
import androidx.activity.result.ActivityResultLauncher
import androidx.health.connect.client.HealthConnectClient
import androidx.health.connect.client.PermissionController
import androidx.health.connect.client.permission.HealthPermission
import androidx.health.connect.client.records.ExerciseSessionRecord
import androidx.health.connect.client.records.metadata.Device
import androidx.health.connect.client.records.metadata.Metadata
import androidx.lifecycle.lifecycleScope
import kotlinx.coroutines.launch
import java.time.Instant
import java.time.ZoneId

/**
 * Writes a finished workout to Health Connect. The site opens
 * `intent://health?start=…&end=…&kind=…&title=…#Intent;scheme=morning30;…`
 * (assets/app-bridge.js) when a workout ends; times are Unix milliseconds.
 * The app only writes exercise sessions and never reads health data. The first
 * time, Health Connect asks for permission.
 */
class HealthActivity : ComponentActivity() {
    private val permissions = setOf(HealthPermission.getWritePermission(ExerciseSessionRecord::class))
    private var workout: ExerciseSessionRecord? = null
    private lateinit var request: ActivityResultLauncher<Set<String>>

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        request = registerForActivityResult(PermissionController.createRequestPermissionResultContract()) { granted ->
            if (granted.containsAll(permissions)) save() else finish()
        }
        workout = intent?.data?.let(::parse)
        if (workout == null || HealthConnectClient.getSdkStatus(this) != HealthConnectClient.SDK_AVAILABLE) {
            finish()
            return
        }
        lifecycleScope.launch {
            val granted = HealthConnectClient.getOrCreate(this@HealthActivity).permissionController.getGrantedPermissions()
            if (granted.containsAll(permissions)) save() else request.launch(permissions)
        }
    }

    private fun save() {
        val record = workout ?: return finish()
        lifecycleScope.launch {
            try {
                HealthConnectClient.getOrCreate(this@HealthActivity).insertRecords(listOf(record))
            } catch (e: Exception) {
                Log.w("morning30", "saving the workout failed", e)
            }
            finish()
        }
    }

    private fun parse(uri: Uri): ExerciseSessionRecord? {
        val start = uri.getQueryParameter("start")?.toLongOrNull() ?: return null
        val end = uri.getQueryParameter("end")?.toLongOrNull() ?: return null
        if (end <= start || end - start > 6 * 3600 * 1000L) return null
        val zone = ZoneId.systemDefault().rules
        val startTime = Instant.ofEpochMilli(start)
        val endTime = Instant.ofEpochMilli(end)
        val type = if (uri.getQueryParameter("kind") == "stretch") {
            ExerciseSessionRecord.EXERCISE_TYPE_STRETCHING
        } else {
            ExerciseSessionRecord.EXERCISE_TYPE_CALISTHENICS
        }
        return ExerciseSessionRecord(
            startTime = startTime,
            startZoneOffset = zone.getOffset(startTime),
            endTime = endTime,
            endZoneOffset = zone.getOffset(endTime),
            // The same workout sent twice replaces itself instead of doubling.
            metadata = Metadata.activelyRecorded(Device(type = Device.TYPE_PHONE), clientRecordId = "morning30-$start"),
            exerciseType = type,
            title = uri.getQueryParameter("title")?.take(80)?.ifBlank { null },
        )
    }
}
