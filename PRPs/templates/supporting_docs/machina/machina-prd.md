# Machina Overview
---
**Machina** is DevQ.ai's unified MCP (Model Context Protocol) Registry and Management Platform that orchestrates 46 MCP servers across the DevQ.ai ecosystem. Machina solves the critical problem of MCP server fragmentation by providing a centralized registry, health monitoring, service discovery, and configuration management system. 

The platform enables DevQ.ai projects to reliably access AI tools and services through both MCP protocol and REST API interfaces, ensuring high availability and performance for AI-powered development workflows. Machina transforms DevQ.ai from having 9 isolated MCP servers to a production-ready ecosystem of 46 interconnected services.

**Target Users:**

- DevQ.ai development teams requiring AI tool integration
- AI agents (Claude, Zed) needing reliable MCP access
- Platform administrators managing service health and configuration
- External DevQ.ai projects consuming MCP services via API

**Value Proposition:**

- **Reliability**: 99.9% uptime through health monitoring and failover
- **Performance**: Load balancing and caching for sub-100ms response times  
- **Consistency**: Standardized API patterns across all 46 MCP servers
- **Scalability**: Docker/Kubernetes deployment supporting unlimited concurrent users

### Core Features  

### 1. **Unified MCP Registry**
- **Purpose**: Central discovery and management of all 46 MCP servers
- **Functionality**: Automatic service discovery, registration, and metadata management
- **Value**: Eliminates manual service tracking and reduces integration complexity
- **Implementation**: FastAPI service with database persistence and real-time updates

### 2. **Service Health Monitoring**
- **Purpose**: Ensure 99.9% uptime across all MCP services
- **Functionality**: Continuous health checks, automatic failover, performance metrics
- **Value**: Prevents service disruptions and enables proactive maintenance
- **Implementation**: Background monitoring with 30-second intervals and Logfire integration

### 3. **Configuration Management CLI**
- **Purpose**: Web interface for managing MCP server configurations
- **Functionality**: Set REQUIRED (TRUE/FALSE), BUILD_PRIORITY (HIGH/MEDIUM/LOW), enable/disable services
- **Value**: Non-technical users can manage complex MCP infrastructure
- **Implementation**: Single-page Python web app with real-time configuration updates

### 4. **Dual Protocol Support**
- **Purpose**: Serve both AI agents (MCP protocol) and DevQ.ai projects (REST API)
- **Functionality**: Simultaneous MCP stdio/WebSocket and HTTP endpoint access
- **Value**: Maximum compatibility and integration flexibility
- **Implementation**: FastAPI + FastMCP hybrid architecture

### 5. **JSON Status Reporting**
- **Purpose**: Real-time status monitoring and integration capabilities
- **Functionality**: JSON endpoints for service status, health metrics, and build progress
- **Value**: Enables external monitoring and automated workflows
- **Implementation**: RESTful JSON APIs with caching and filtering

### 6. **Build Type Classification**
- **Purpose**: Organize servers by implementation approach for development efficiency
- **Functionality**: List servers by FastMCP, FastAPI, External, and Conversion categories
- **Value**: Streamlines development planning and resource allocation
- **Implementation**: Database categorization with filtering and reporting views


### User Personas

#### **1. DevQ.ai Developer**
- **Goals**: Quickly integrate MCP tools into projects, troubleshoot service issues
- **Journey**: Browse available services → Configure for project → Monitor health → Debug issues
- **Pain Points**: Complex MCP setup, service downtime, unclear documentation

#### **2. Platform Administrator** 
- **Goals**: Maintain service health, manage configurations, monitor performance
- **Journey**: Review dashboard → Adjust priorities → Deploy updates → Monitor metrics
- **Pain Points**: Manual health checks, configuration drift, scaling complexity

#### **3. AI Agent (Claude/Zed)**
- **Goals**: Reliable access to MCP tools with fast response times
- **Journey**: Connect to registry → Discover tools → Execute commands → Handle errors
- **Pain Points**: Service timeouts, inconsistent APIs, connection failures

### Key User Flows

#### **Service Discovery Flow**
1. User accesses Machina dashboard at `https://machina.devq.ai`
2. Browse 46 available services categorized by build type
3. Filter by REQUIRED status, BUILD_PRIORITY, or functionality
4. View service details, health status, and integration instructions
5. Copy API endpoints or MCP connection strings for integration

#### **Configuration Management Flow**
1. Administrator logs into configuration panel
2. Select service from categorized list (FastMCP/FastAPI/External/Conversion)
3. Adjust REQUIRED flag and BUILD_PRIORITY level
4. Save changes with automatic validation
5. Real-time updates propagate to all consumers

#### **Health Monitoring Flow**
1. System performs automated health checks every 30 seconds
2. Failures trigger immediate failover to healthy instances
3. Administrators receive alerts for persistent issues
4. Dashboard displays real-time health status with historical trends
5. JSON status API provides programmatic access to metrics

### UI/UX Considerations

#### **Single-Page Configuration Interface**
- **Layout**: Three-column design (Service List | Configuration Panel | Status Monitor)
- **Interactions**: Click to select, toggle switches for boolean flags, dropdown for priorities
- **Responsiveness**: Desktop-first design with mobile compatibility
- **Performance**: Real-time updates without page refreshes

#### **Service Categorization Display**
- **Visual Design**: Color-coded cards by build type (FastMCP=Blue, FastAPI=Green, External=Orange, Conversion=Purple)
- **Information Hierarchy**: Service name → Status badge → Priority indicator → Description
- **Filtering**: Multi-select checkboxes for build type, status, and priority
- **Search**: Real-time text search across service names and descriptions

### Technical Architecture  

### System Components

#### **1. Core Registry Service** (`machina-registry/`)
```python
# FastAPI + FastMCP hybrid service
- Primary Service: FastAPI application with dual protocol support
- MCP Integration: FastMCP component for native MCP protocol handling
- Database: PostgreSQL for service metadata and configuration persistence
- Cache Layer: Redis for high-performance service discovery and health status
- Message Queue: Redis pub/sub for real-time configuration updates
```

#### **2. Service Discovery Engine** (`discovery/`)
```python
# Automatic detection and registration of MCP services
- Local Scanner: Detects services in /mcp/mcp-servers/ directory
- External Integrator: Connects to third-party MCP servers (Stripe, PayPal, etc.)
- Project Converter: Identifies FastAPI apps ready for MCP conversion
- Docker Monitor: Discovers containerized MCP services
- Health Prober: Validates service functionality during discovery
```

#### **3. Configuration Management UI** (`ui/`)
```python
# Single-page Python web application
- Framework: FastAPI with Jinja2 templates and HTMX for interactivity
- State Management: Real-time updates via WebSocket connections
- Form Handling: JSON validation with Pydantic models
- Responsive Design: Tailwind CSS with mobile-first approach
- Security: Bearer token authentication with role-based access
```

#### **4. Health Monitoring System** (`monitoring/`)
```python
# Background service for continuous health checking
- Health Checker: Async monitoring with 30-second intervals
- Failover Manager: Automatic routing to healthy instances
- Metrics Collector: Performance data aggregation with Logfire
- Alert System: Notification dispatch for critical failures
- Load Balancer: Request distribution across healthy instances
```

#### **5. Docker Orchestration** (`docker/`)
```yaml
# Multi-container deployment with service mesh
services:
  machina-registry:    # Core registry service
  machina-ui:          # Configuration interface
  machina-monitor:     # Health monitoring
  postgres:            # Configuration database
  redis:               # Cache and message queue
  nginx:               # Load balancer and SSL termination
```

### Data Models

#### **Service Registration Schema**
```python
class ServiceRegistration(BaseModel):
    name: str
    build_type: Literal["fastmcp", "fastapi", "external", "conversion"]
    required: bool = False
    build_priority: Literal["HIGH", "MEDIUM", "LOW"] = "MEDIUM"
    status: Literal["healthy", "unhealthy", "unknown"] = "unknown"
    location: str  # File path, URL, or container name
    protocol: Literal["stdio", "http", "websocket"]
    description: str
    tags: List[str] = []
    config: Dict[str, Any] = {}
    health_endpoint: Optional[str] = None
    api_key: Optional[str] = None
    created_at: datetime
    updated_at: datetime
```

#### **Configuration Schema**
```python
class ServiceConfig(BaseModel):
    service_name: str
    required: bool
    build_priority: Literal["HIGH", "MEDIUM", "LOW"]
    enabled: bool = True
    environment_vars: Dict[str, str] = {}
    scaling_config: Dict[str, Any] = {}
    updated_by: str
    updated_at: datetime
```

#### **Health Status Schema**
```python
class HealthStatus(BaseModel):
    service_name: str
    status: Literal["healthy", "unhealthy", "unknown"]
    last_check: datetime
    response_time_ms: float
    error_message: Optional[str] = None
    consecutive_failures: int = 0
    uptime_percentage: float
```

### APIs and Integrations

#### **REST API Endpoints**
```python
# Service Discovery
GET /api/services                    # List all services
GET /api/services/{name}            # Get service details
GET /api/services/by-type/{type}    # Filter by build type

# Configuration Management  
PUT /api/services/{name}/config     # Update service configuration
GET /api/services/{name}/config     # Get current configuration
POST /api/services/{name}/toggle    # Enable/disable service

# Health Monitoring
GET /api/health                     # Overall system health
GET /api/health/{name}              # Individual service health
GET /api/health/status.json         # JSON status for external monitoring

# Build Management
GET /api/build/priority/{level}     # Services by priority level
GET /api/build/required             # Required services only
POST /api/build/deploy/{name}       # Deploy specific service
```

#### **MCP Protocol Integration**
```python
# Native MCP tool exposure
@mcp.tool
async def list_available_services():
    """List all registered MCP services"""

@mcp.tool  
async def call_service_tool(service: str, tool: str, **kwargs):
    """Route tool calls to specific MCP services"""

@mcp.resource
async def service_documentation(service: str):
    """Get documentation for specific service"""
```

#### **WebSocket Real-time Updates**
```python
# Configuration changes
ws://machina.devq.ai/ws/config      # Real-time config updates
ws://machina.devq.ai/ws/health      # Live health status
ws://machina.devq.ai/ws/builds      # Build progress monitoring
```

### Infrastructure Requirements

#### **Production Deployment**

- **Container Platform**: Docker with Kubernetes orchestration
- **Compute**: 3 nodes × 4 CPU, 16GB RAM for high availability
- **Storage**: 100GB SSD for database, 50GB for logs and cache
- **Network**: Load balancer with SSL termination and health checks
- **Monitoring**: Prometheus + Grafana + Logfire integration

#### **Development Environment**

- **Local Setup**: Docker Compose with hot reloading
- **Database**: PostgreSQL 15 with automatic migrations
- **Cache**: Redis 7 with persistence enabled
- **Reverse Proxy**: Nginx with development SSL certificates

#### **Security Requirements**

- **Authentication**: Bearer tokens with 24-hour expiration
- **Authorization**: Role-based access (admin, developer, viewer)
- **Encryption**: TLS 1.3 for all external communications
- **Secrets**: Kubernetes secrets for API keys and database credentials
- **Audit**: Complete audit trail for all configuration changes

### Development Roadmap  

### **Phase 1: MVP Core Registry (Week 1-2)**
**Scope**: Basic service discovery and registration with manual configuration

**Deliverables**:

- Core FastAPI application with PostgreSQL integration
- Service discovery for existing 9 MCP servers in `/mcp/mcp-servers/`
- Basic REST API for service listing and health checks
- Manual service registration via API endpoints
- Docker Compose development environment
- Unit tests for core functionality

**Success Criteria**:

- All 9 existing MCP servers automatically discovered and registered
- REST API returns accurate service status and metadata
- Basic health checking confirms service availability
- Docker environment runs without manual intervention

### **Phase 2: Configuration Management UI (Week 3)**
**Scope**: Single-page web interface for service configuration

**Deliverables**:

- FastAPI web application with Jinja2 templates
- Configuration form with REQUIRED and BUILD_PRIORITY fields
- Real-time updates via HTMX for seamless user experience
- Service filtering by build type (FastMCP, FastAPI, External, Conversion)
- Configuration persistence with validation
- WebSocket integration for live updates

**Success Criteria**:

- Non-technical users can modify service configurations without API calls
- Changes reflect immediately in the interface without page refresh
- Form validation prevents invalid configurations
- All 46 target services display with correct categorization

### **Phase 3: External Service Integration (Week 4)**
**Scope**: Integration with 8 external/third-party MCP servers

**Deliverables**:

- External service discovery for official MCP servers (Redis, SQLite, etc.)
- API integration with service providers (Stripe, PayPal, Shopify)
- Configuration templates for external service setup
- Health monitoring for external dependencies
- Fallback handling for external service failures
- Documentation for external service onboarding

**Success Criteria**:

- 8 external services successfully integrated and monitored
- External service failures don't impact core registry functionality
- Clear documentation enables easy addition of new external services
- Health checks accurately reflect external service status

### **Phase 4: Project Conversion System (Week 5)**
**Scope**: Convert existing DevQ.ai projects (Bayes, Darwin) to MCP servers

**Deliverables**:

- Conversion framework for FastAPI → MCP server transformation
- Bayes project conversion to bayes-mcp server
- Darwin project conversion to darwin-mcp server
- Automated deployment pipeline for converted services
- Testing framework for converted services
- Migration documentation and best practices

**Success Criteria**:

- Bayes and Darwin projects successfully serve as MCP servers
- Converted services maintain original functionality while adding MCP capabilities
- Automated conversion process documented for future projects
- Performance benchmarks meet or exceed original implementations

### **Phase 5: Advanced Monitoring & Health Management (Week 6)**
**Scope**: Production-ready monitoring with automatic failover

**Deliverables**:

- Background health monitoring service with 30-second intervals
- Automatic failover and load balancing for multiple service instances
- Logfire integration for comprehensive observability
- Performance metrics collection and historical trending
- Alert system for critical service failures
- Dashboard for real-time system health visualization

**Success Criteria**:

- System automatically recovers from individual service failures
- Performance metrics accurately reflect system load and response times
- Alert system provides actionable notifications without false positives
- Dashboard provides clear visibility into overall system health

### **Phase 6: Production Deployment & Scaling (Week 7-8)**
**Scope**: Kubernetes deployment with full 46-service ecosystem

**Deliverables**:

- Kubernetes manifests for production deployment
- Helm charts for easy deployment management
- Horizontal pod autoscaling based on load metrics
- Complete implementation of remaining 16 FastMCP services
- Implementation of 11 complex FastAPI services
- Production monitoring and logging infrastructure
- Disaster recovery procedures and documentation

**Success Criteria**:

- All 46 MCP services running in production environment
- System handles 1000+ concurrent requests without degradation
- Automatic scaling responds appropriately to load changes
- Complete documentation enables team maintenance and expansion

### Logical Dependency Chain

### **Foundation Layer (Must Build First)**
1. **Core Registry Database Schema** - All other components depend on service metadata
2. **Service Discovery Engine** - Required before any services can be registered
3. **Basic REST API** - Needed for all other components to interact with registry
4. **Docker Development Environment** - Essential for consistent development experience

### **Service Integration Layer (Build Second)**
1. **Health Monitoring System** - Required before external dependencies
2. **Existing MCP Server Integration** - Establishes patterns for other integrations
3. **Configuration Persistence** - Needed before UI development
4. **Basic Load Balancing** - Required for multi-instance deployments

### **User Interface Layer (Build Third)**
1. **Configuration Management UI** - Depends on REST API and configuration persistence
2. **Real-time Update System** - Requires WebSocket infrastructure and UI foundation
3. **Service Filtering and Search** - Builds upon basic UI with enhanced functionality
4. **Dashboard and Monitoring Views** - Requires health monitoring and UI framework

### **External Integration Layer (Build Fourth)**
1. **External Service Discovery** - Depends on core discovery patterns
2. **Third-party API Integration** - Builds upon external discovery framework
3. **Conversion Framework** - Requires understanding of service patterns from previous phases
4. **Advanced Health Monitoring** - Extends basic monitoring to external dependencies

### **Production Readiness Layer (Build Last)**
1. **Kubernetes Deployment** - Requires all core functionality to be stable
2. **Advanced Monitoring and Alerting** - Builds upon health monitoring foundation
3. **Performance Optimization** - Requires load testing with full service suite
4. **Documentation and Training** - Depends on stable, feature-complete system

### **Atomic Feature Development Approach**

#### **Service Registration Module**

- **Atomic Unit**: Single service can be registered, retrieved, and updated
- **Build Upon**: Add bulk operations, filtering, and advanced querying
- **Improvement Path**: Enhance with validation, versioning, and audit trails

#### **Health Monitoring Module**

- **Atomic Unit**: Single service health check with basic status reporting
- **Build Upon**: Add multiple instances, failover logic, and performance metrics
- **Improvement Path**: Enhance with predictive monitoring and automated remediation

#### **Configuration Management Module**

- **Atomic Unit**: Single configuration parameter can be modified and persisted
- **Build Upon**: Add bulk updates, configuration templates, and change management
- **Improvement Path**: Enhance with approval workflows and configuration validation

### Risks and Mitigations  

### **Technical Challenges**

#### **Risk**: Service Discovery Complexity

- **Description**: Automatically discovering 46 diverse MCP servers across different implementation patterns
- **Impact**: Core functionality failure, incomplete service registry
- **Mitigation**: 

  - Implement discovery in phases starting with known patterns
  - Create standardized service metadata format
  - Build fallback manual registration system
  - Extensive testing with mock services

#### **Risk**: Health Monitoring Accuracy

- **Description**: False positives/negatives in health checks leading to unnecessary failovers
- **Impact**: Service instability, poor user experience, operational overhead
- **Mitigation**:

  - Implement multiple health check methods (endpoint, heartbeat, functional test)
  - Use configurable thresholds and consecutive failure requirements
  - Add manual override capabilities for administrators
  - Comprehensive logging for health check debugging

#### **Risk**: Performance Under Load

- **Description**: Registry becomes bottleneck when serving 46 services to multiple consumers
- **Impact**: Slow response times, service timeouts, system unavailability
- **Mitigation**:

  - Implement caching layer with Redis for frequently accessed data
  - Use async/await patterns throughout for non-blocking operations
  - Add load testing early in development cycle
  - Plan for horizontal scaling with Kubernetes

### **MVP Definition and Scope Management**

#### **Risk**: Feature Scope Creep

- **Description**: Adding advanced features before core functionality is stable
- **Impact**: Delayed MVP delivery, complex debugging, unstable foundation
- **Mitigation**:

  - Strict adherence to phase-based development roadmap
  - MVP success criteria clearly defined and measurable
  - Regular scope review meetings with stakeholder approval required for changes
  - Feature flag system to disable advanced functionality during development

#### **Risk**: Integration Complexity

- **Description**: 46 different MCP servers with varying implementation patterns and requirements
- **Impact**: Extended development time, maintenance burden, user confusion
- **Mitigation**:

  - Start with simplest services and establish patterns
  - Create standardized adapter interfaces for different service types
  - Build comprehensive service categorization system
  - Prioritize high-value, low-complexity services for early wins

### **Resource and Operational Constraints**

#### **Risk**: External Service Dependencies

- **Description**: Third-party MCP servers (Stripe, PayPal, etc.) outside DevQ.ai control
- **Impact**: Service unavailability, authentication failures, API changes
- **Mitigation**:

  - Implement circuit breaker pattern for external service calls
  - Cache external service responses where appropriate
  - Build comprehensive error handling and user feedback
  - Maintain backup plans for critical external services

#### **Risk**: Configuration Management Complexity

- **Description**: Managing configuration for 46 services across multiple environments
- **Impact**: Configuration drift, deployment failures, security vulnerabilities
- **Mitigation**:

  - Implement configuration as code with version control
  - Use environment-specific configuration files with validation
  - Build configuration comparison and drift detection tools
  - Establish clear configuration change approval process

#### **Risk**: DevOps and Deployment Complexity

- **Description**: Coordinating deployment of 46+ services in production environment
- **Impact**: Deployment failures, service incompatibilities, difficult rollbacks
- **Mitigation**:

  - Use Docker containers for consistent deployment environments
  - Implement gradual rollout with blue-green deployment strategies
  - Build comprehensive automated testing pipeline
  - Create detailed runbooks for common operational scenarios

### Appendix  

**Research Findings**

#### **MCP Server Analysis Results**
Based on comprehensive analysis of DevQ.ai's current infrastructure:

- **9 servers already implemented** (19.6% completion rate)
- **2 existing projects ready for conversion** (Bayes, Darwin)
- **8 external services available for integration** (official providers)
- **27 services require full implementation** from scratch
- **Strong FastAPI foundation** with 19 existing applications identified

#### **Technology Stack Validation**

- **FastAPI + FastMCP Hybrid**: Confirmed as optimal approach for dual protocol support
- **PostgreSQL**: Suitable for service metadata and configuration persistence
- **Redis**: Proven effective for caching and real-time updates in DevQ.ai projects
- **Docker + Kubernetes**: Aligns with DevQ.ai infrastructure standards
- **Logfire Integration**: Mandatory for observability per DevQ.ai standards

#### **Performance Requirements**

- **Response Time Target**: <100ms for service discovery and health checks
- **Throughput Target**: 1000+ concurrent requests without degradation
- **Availability Target**: 99.9% uptime (8.77 hours downtime per year)
- **Scalability Target**: Support for 100+ MCP services in future expansions

### **Technical Specifications**

#### **Database Schema**
```sql
-- Core service registry table
CREATE TABLE services (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) UNIQUE NOT NULL,
    build_type VARCHAR(50) NOT NULL CHECK (build_type IN ('fastmcp', 'fastapi', 'external', 'conversion')),
    required BOOLEAN DEFAULT FALSE,
    build_priority VARCHAR(10) NOT NULL CHECK (build_priority IN ('HIGH', 'MEDIUM', 'LOW')),
    status VARCHAR(20) DEFAULT 'unknown',
    location TEXT NOT NULL,
    protocol VARCHAR(20) NOT NULL,
    description TEXT,
    config JSONB DEFAULT '{}',
    tags TEXT[] DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Health monitoring table
CREATE TABLE health_checks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_id UUID REFERENCES services(id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL,
    response_time_ms FLOAT,
    error_message TEXT,
    checked_at TIMESTAMPTZ DEFAULT NOW()
);

-- Configuration management table
CREATE TABLE service_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_id UUID REFERENCES services(id) ON DELETE CASCADE,
    config_key VARCHAR(255) NOT NULL,
    config_value TEXT NOT NULL,
    updated_by VARCHAR(255),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(service_id, config_key)
);
```

#### **Docker Compose Development Environment**
```yaml
version: '3.8'
services:
  machina-registry:
    build: .
    ports: ["8000:8000"]
    environment:
      - DATABASE_URL=postgresql://postgres:password@postgres:5432/machina
      - REDIS_URL=redis://redis:6379
      - LOGFIRE_TOKEN=${LOGFIRE_TOKEN}
    depends_on: [postgres, redis]
    volumes: ["./src:/app/src:ro"]

  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: machina
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
    volumes: ["postgres_data:/var/lib/postgresql/data"]

  redis:
    image: redis:7-alpine
    volumes: ["redis_data:/data"]

  nginx:
    image: nginx:alpine
    ports: ["80:80", "443:443"]
    volumes: ["./nginx.conf:/etc/nginx/nginx.conf:ro"]
    depends_on: [machina-registry]

volumes:
  postgres_data:
  redis_data:
```
--- 

apiVersion: v1
kind: Service
metadata:
  name: machina-registry-service
spec:
  selector:
    app: machina-registry
  ports:
  - port: 80
    targetPort: 8000
  type: LoadBalancer
```

#### **API Documentation Standards**

- **OpenAPI 3.0**: Full specification with examples and validation schemas
- **Authentication**: Bearer token requirement for all configuration endpoints
- **Rate Limiting**: 1000 requests per minute per API key
- **Versioning**: URL path versioning (/api/v1/) with backward compatibility
- **Error Handling**: Consistent JSON error responses with detailed messages
