from pydantic_ai import Agent
import logfire
from typing import Dict, Any

@logfire.instrument()
class DevQaiAgent:
    """Standard DevQ.ai agent with Logfire integration."""

    def __init__(self, model: str = 'claude-3-7-sonnet-20250219'):
        self.agent = Agent(
            model,
            system_prompt="""You are a DevQ.ai agent.
            Requirements:
            - Use FastAPI for all web services
            - Include PyTest for all functionality
            - Log everything through Logfire
        )

    @logfire.instrument()
    async def process_task(self, task: Dict[str, Any]) -> Dict[str, Any]:
        """Process task with full observability."""
        with logfire.span("Agent task processing", task_id=task.get('id')):
            result = await self.agent.run(task['prompt'])
            logfire.info("Task completed", task_id=task.get('id'), result_length=len(str(result)))
            return result
