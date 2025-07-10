
## Inputs
### {project}=name
### {local_repo}="/Users/dionedge/devqai/"{project}
### {remote_repo}="https://github.com/devqai/"{project}

---

## File Inventory
### source_filepath: [primer.md]((/Users/dionedge/devqai/devgen/primer.md)
#### destination_filepath: [primer.yaml](/Users/dionedge/devqai/{project}/primer.yaml)
#### transformation 
- complete and save ./devqai/primer.md
- script will generate ./devqai/{project}/primer.yaml
- script will generate a cleam ./devqai/primer.md
- dart script will generate ./devqai/{project}/context_year-mm-dd.md

### source_filepath: [./devqai/devgen/rules/](./devqai/devgen/rules/)
#### destination_filepath: [.rules](/Users/dionedge/devqai/{project}) 
#### transformation 
- Generate file with selected rules
- Move .rules to {project} root

### source_filepath: [./devqai/.env](/Users/dionedge/devqai/.env)
#### destination_filepath: [./devqai/{project}.env](/Users/dionedge/devqai/{project}/.env)
#### transformation: copy source file to {project} root

### source_filepath: [./CHANGELOG.md](/Users/dionedge/devqai/CHANGELOG.LOG.md) 
#### destination_filepath: [./CHANGELOG.md](/Users/dionedge/devqai/{project}/CHANGELOG.md)
#### transformation: Create empty CHANGELOG in project root

---

## Rules

### [A. Development Environment](/Users/dionedge/devqai/devgen/rules/dev_environment.md)
#### Attach entire file contents to `.rules`
#### **[ ] Load Rules**

### [B. Backend Development Rules](/Users/dionedge/devqai/devgen/rules/backend_rules.md)
#### Attach entire file contents to `.rules`
#### **[ ] Load Rules**

### [C. Frontend Development Rules](/Users/dionedge/devqai/devgen/rules/frontend_rules.md)
#### Attach entire file contents to `.rules`
#### **[ ] Load Rules**

### [D. Git Workflow Standards](/Users/dionedge/devqai/devgen/rules/git_workflow.md)
#### Attach entire file contents to `.rules`
#### **[ ] Load Rules**

### [E. Common Development Rules](/Users/dionedge/devqai/devgen/rules/common_rules.md)
#### Attach entire file contents to `.rules`
#### **[ ] Load Rules**

### [F. Frameworks](/Users/dionedge/devqai/devgen/rules/frameworks.md)
#### Attach selected sections to `.rules`
#### Selected
## 1. Agent Frameworks
- [ ] [Archon](https://github.com/coleam00/Archon)
- [ ] [AgentSeek](https://github.com/Fosowl/agenticSeek)
- [ ] [Binome](https://github.com/binome-dev/graphite/)
- [ ] [Flowise](https://github.com/FlowiseAI/FlowiseDocs)
- [ ] [Local-operator](https://github.com/damianvtran/local-operator)
- [ ] [Pydantic-ai](https://github.com/pydantic/pydantic-ai)
## 2. Required Back End
- [ ] Auth: [Better-auth](https://github.com/better-auth/better-auth/)
- [ ] Database_Migration: [Alembic](https://github.com/sqlalchemy/alembic)
- [ ] DB_ToolKit: [SQLAlchemy](https://github.com/sqlalchemy/sqlalchemy)
- [ ] Deployment: [Coolify](https://github.com/coollabsio/coolify/)
- [ ] Env_Variables: [Python-dotenv](https://pypi.org/project/python-dotenv/)
- [ ] GitHub_Code_Review: [CodeRabbit](https://github.com/coderabbitai/coderabbit-docs/)
- [ ] Logging: [logfire](https://github.com/pydantic/logfire)
- [ ] Testing: [Pytest](https://pypi.org/project/pytest/)
## 3. Other Back End
- [ ] Tools: [ACI_Tools](https://github.com/aipotheosis-labs/aci)
- [ ] Web_Framewok: [FastAPI](https://github.com/fastapi/fastapi)
- [ ] Web_Scraping: [crawl4ai](https://github.com/unclecode/crawl4ai)
- [ ] Zero_Knowledge_Proof: [union](https://github.com/unionlabs/union)
## 4. Computational Frameworks
- [ ] [Game-Theory](https://github.com/Axelrod-Python/Axelrod)
- [ ] [Genetic-Algorithm](https://github.com/ahmedfgad/GeneticAlgorithmPython/)
- [ ] [NumPy](https://github.com/numpy/numpy)
- [ ] [Pandas](https://github.com/pandas-dev/pandas)
- [ ] [PyMC](https://github.com/pymc-devs/pymc)
- [ ] [PyTorch](https://github.com/pytorch/pytorch)
- [ ] [Random-Forest](https://github.com/pyensemble/wildwood)
- [ ] [SciComPy](https://github.com/yoavram/SciComPy)
## 5. Database
- [ ] [graphiti](https://github.com/getzep/graphiti)
- [ ] [logfire](https://logfire.pydantic.dev/docs/)
- [ ] [neo4j](https://github.com/neo4j)
- [ ] [surrealdb](https://github.com/surrealdb/surrealdb)
## 6. Required Front End
- [ ] [Anime.js](https://github.com/juliangarnier/anime/)
- [ ] [Next.js](https://github.com/nextjs)
- [ ] [SaaS App](https://github.com/adrianhajdin/saas-app/)
- [ ] [Shadcn UI](https://github.com/birobirobiro/awesome-shadcn-ui/)
- [ ] [Tailwind CSS](https://github.com/tailwindlabs/tailwindcss)
- [ ] [Tiptap](https://github.com/ueberdosis/tiptap)
- [ ] [Tweakcn](https://github.com/jnsahaj/tweakcn/)
## 7. Other Front End 
- [ ] [Bokeh](https://github.com/bokeh/bokeh)
- [ ] [D3.js](https://github.com/d3/d3)
- [ ] [docusaurus](https://docusaurus.io/)
- [ ] [Flowgram.ai](https://github.com/bytedance/flowgram.ai)
- [ ] [Panel](https://github.com/holoviz/panel)
- [ ] [Streamlit](https://github.com/streamlit)
- [ ] [Tersa](https://github.com/haydenbleasel/tersa)
## 8. Finance Frameworks
- [ ] [accounting](https://github.com/ekmungai/python-accounting)
- [ ] [ai-hedge-fund](https://github.com/virattt/ai-hedge-fund)
- [ ] [financialdatasets](https://www.financialdatasets.ai/)
- [ ] [kalshi](https://github.com/Kalshi/kalshi-starter-code-python)
- [ ] [plaid](https://plaid.com/docs/)
- [ ] [qlib](https://github.com/microsoft/qlib)
- [ ] [Stripe](https://docs.stripe.com/)
- [ ] [tokencost](https://github.com/AgentOps-AI/tokencost)
- [ ] [xero](https://github.com/XeroAPI/xero-python)
## 9. Data Model
- [ ] [agentql](https://github.com/tinyfish-io/agentql)
- [ ] [cube](https://cube.dev/docs/product/introduction/)
- [ ] [dbt](https://docs.getdbt.com/)
- [ ] [linkml](https://github.com/linkml/)
- [ ] [lookml](https://cloud.google.com/looker/docs/reference/lookml-quick-reference)
- [ ] [omni](https://docs.omni.co/docs)
- [ ] [sqlglot](https://github.com/tobymao/sqlglot)
- [ ] [sqlx](https://cloud.google.com/dataform/docs/)

### [G. Ptolemies Knowledge Base](/Users/dionedge/devqai/devgen/rules/knowledge_base.md)
#### Attach selected sections to `.rules`
#### **[ ] Load Rules**
#### 1. Loaded Knowledge Base Documention
- [x] [fastapi](https://fastapi.tiangolo.com/tutorial/metadata/)
- [x] [graphiti](https://github.com/getzep/graphiti)
- [x] [logfire](https://logfire.pydantic.dev/docs/)
- [x] [surrealdb](https://github.com/surrealdb/surrealdb)

### [H. Cloud Engineering Rules](/Users/dionedge/devqai/devgen/rules/cloud_rules.md)
#### Attach selected sections to `.rules`
#### **[ ] Load Rules**

### [I. Design System](/Users/dionedge/devqai/devgen/rules/design_system.md)
#### Attach selected sections to `.rules`
#### Selected
- [ ] ## 1. Status Colors
- [ ] ## 2. Priority Colors
- [ ] ## 3. UI Elements
- [ ] ## 4. Midnight UI (Elegant & Minimal)
- [ ] ## 5. Cyber Dark (Futuristic & High Contrast)
- [ ] ## 6. Electric Dream (Vibrant & Edgy)
- [ ] ## 7. Modern Soft (Neutral & Minimal)
- [ ] ## 8. Pastel UI (Soft & Friendly)
- [ ] ## 9. Cyber Cotton Candy (Soft but Electric)

---

### Append to [.rules](/Users/dionedge/devqai/{project}/.rules)
- A. Development Environment
- B. Backend Development Rules
- C. Frontend Development Rules
- D. Git Workflow Standards
- E. Common Development Rules
- F. Frameworks
- G. Ptolemies Knowledge Base
- H. Cloud Engineering Rules
- I. Design System

---

## Dev End Settings Files
### [Dev Report](/Users/dionedge/devqai/dev_report.md)
### [./mcp/mcp-servers.json](/Users/dionedge/devqai/mcp/mcp-servers.json)
### [.claude/local.settings.json](/Users/dionedge/devqai/.claude/local.settings.json)
### [.zed/settings.json](/Users/dionedge/devqai/.zed/settings.json)
### [.zed/zed-terminal-config.zsh](/Users/dionedge/devqai/.zed/zed-terminal-config.zsh)

---

**[ ] GENERATE ./devqai/{project}/primer.yaml**