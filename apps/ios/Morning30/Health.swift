import HealthKit

/// Writes finished workouts to Apple Health. The app only writes workouts and
/// never reads health data. The permission sheet appears the first time a
/// workout is saved; the person can change it later in the Health app.
enum Health {
    private static let store = HKHealthStore()

    /// Saves one workout. `kind` is "stretch" or "strength".
    static func saveWorkout(start: Date, end: Date, title: String, kind: String) {
        guard HKHealthStore.isHealthDataAvailable(), end > start else { return }
        let type = HKObjectType.workoutType()
        store.requestAuthorization(toShare: [type], read: []) { _, _ in
            guard store.authorizationStatus(for: type) == .sharingAuthorized else { return }
            let configuration = HKWorkoutConfiguration()
            configuration.activityType = kind == "stretch" ? .flexibility : .functionalStrengthTraining
            configuration.locationType = .indoor
            let builder = HKWorkoutBuilder(healthStore: store, configuration: configuration, device: .local())
            builder.beginCollection(withStart: start) { began, _ in
                guard began else { return }
                var metadata: [String: Any] = [HKMetadataKeyIndoorWorkout: true]
                if !title.isEmpty {
                    metadata[HKMetadataKeyWorkoutBrandName] = title
                }
                builder.addMetadata(metadata) { _, _ in
                    builder.endCollection(withEnd: end) { ended, _ in
                        guard ended else { return }
                        builder.finishWorkout { _, error in
                            if let error { NSLog("morning30: saving the workout failed: \(error.localizedDescription)") }
                        }
                    }
                }
            }
        }
    }
}
