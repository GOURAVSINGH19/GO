what is graceful shutdown ?

Imagine you're running a backend server with 3 workers processing customer orders.

Suddenly, your server receives a shutdown signal.

Should it immediately kill all workers?

No. Some workers might be processing important orders.

Instead, we want to:
1. Stop accepting new jobs.
2. Allow existing jobs to finish.
3. Stop all workers safely.
4. Exit the application.
This process is called Graceful Shutdown.



User presses Ctrl+C

Shutdown starts
1. Producer stops submitting new jobs.
2. Producer closes jobs.
3. Workers finish their current jobs.
4. Workers receive and process Jobs 4, 5 and 6.
5. Workers exit after the channel becomes empty.
6. wg.Wait() returns.
   Server stopped gracefully