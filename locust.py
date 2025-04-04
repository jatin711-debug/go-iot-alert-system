from locust import HttpUser, task, between
import random

class APIUser(HttpUser):
    # Simulates user wait time between requests
    wait_time = between(1, 3)

    @task(1)
    def health_check(self):
        """Check if the service is up."""
        with self.client.get("/api/health", catch_response=True) as response:
            if response.status_code != 200:
                response.failure(f"Health check failed with status {response.status_code}")
            else:
                response.success()

    @task(2)
    def get_alerts(self):
        """Test fetching alerts with a random asset_id."""
        asset_id = 1
        with self.client.get(f"/api/alerts?asset_id={asset_id}", catch_response=True) as response:
            if response.status_code != 200:
                response.failure(f"Failed to get alerts for asset_id={asset_id}, status {response.status_code}")
            else:
                response.success()
