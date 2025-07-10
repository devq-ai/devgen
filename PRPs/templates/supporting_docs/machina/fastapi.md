# Required FastAPI app structure with Logfire integration
# api/main.py
from fastapi import FastAPI, HTTPException, Depends
from fastapi.middleware.cors import CORSMiddleware
import logfire
from contextlib import asynccontextmanager

# Configure Logfire before app creation
logfire.configure(token='your_logfire_token')

@asynccontextmanager
async def lifespan(app: FastAPI):
    # Startup
    logfire.info("Starting DevQ.ai API server")
    yield
    # Shutdown
    logfire.info("Shutting down DevQ.ai API server")

# Required FastAPI app configuration
app = FastAPI(
    title="DevQ.ai API",
    description="AI-powered development platform API",
    version="1.0.0",
    lifespan=lifespan
)

# Required middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:3000"],  # Frontend origin
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Required Logfire instrumentation
logfire.instrument_fastapi(app, capture_headers=True)
logfire.instrument_httpx()
logfire.instrument_requests()

# Required health endpoint pattern
@app.get("/health")
async def health_check():
    with logfire.span("health-check"):
        return {"status": "healthy", "service": "devqai-api"}

# Required error handling
@app.exception_handler(Exception)
async def global_exception_handler(request, exc):
    logfire.error("Unhandled exception", exc_info=exc)
    return {"error": "Internal server error"}

# Required route pattern with Logfire spans
@app.get("/api/workflows")
async def get_workflows():
    with logfire.span("get-workflows") as span:
        try:
            # Real implementation required
            workflows = await fetch_workflows_from_db()
            span.set_attribute("workflow_count", len(workflows))
            return workflows
        except Exception as e:
            logfire.error("Failed to fetch workflows", exc_info=e)
            raise HTTPException(status_code=500, detail="Failed to fetch workflows")