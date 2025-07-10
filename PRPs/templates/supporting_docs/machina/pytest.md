# conftest.py - Required test configuration
import pytest
import asyncio
from fastapi.testclient import TestClient
import logfire

# Configure test logging
logfire.configure(send_to_logfire=False)  # Disable in tests

@pytest.fixture(scope="session")
def event_loop():
    loop = asyncio.get_event_loop_policy().new_event_loop()
    yield loop
    loop.close()

@pytest.fixture
def client():
    from main import app
    with TestClient(app) as test_client:
        yield test_client

# Required test pattern
def test_health_endpoint(client):
    """All endpoints must have tests."""
    response = client.get("/health")
    assert response.status_code == 200
    assert response.json()["status"] == "healthy"
