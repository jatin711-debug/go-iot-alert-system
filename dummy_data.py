import psycopg2
from psycopg2.extras import execute_values
from datetime import datetime, timedelta
import random

# Database connection configuration
DB_CONFIG = {
    'dbname': 'alerts',
    'user': 'root',
    'password': 'secret',
    'host': 'localhost',
    'port': '5432'
}

# Possible severity levels
SEVERITY_LEVELS = ['Low', 'Medium', 'High', 'Critical']

# Generate dummy data
def generate_dummy_alerts(num_records=100000000):
    alerts = []
    for _ in range(num_records):
        asset_id = random.randint(1, 50)  # Assuming asset IDs are between 1 and 50
        severity = random.choice(SEVERITY_LEVELS)
        created_at = datetime.now() - timedelta(days=random.randint(0, 30))
        alerts.append((asset_id, severity, created_at))
    return alerts

# Insert dummy data into the database
def insert_dummy_data():
    alerts = generate_dummy_alerts()
    insert_query = """
        INSERT INTO public.alerts (asset_id, severity, created_at)
        VALUES %s
    """
    try:
        with psycopg2.connect(**DB_CONFIG) as conn:
            with conn.cursor() as cursor:
                execute_values(cursor, insert_query, alerts)
                conn.commit()
        print(f"Inserted {len(alerts)} dummy alerts successfully.")
    except Exception as e:
        print("Error inserting data:", e)

if __name__ == "__main__":
    insert_dummy_data()
