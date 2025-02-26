
<h1>Swedish Club Management API</h1>

<p>
  <strong>A task management API built with Go to streamline executive responsibilities within the University of Auckland Swedish Club.</strong>
</p>

<h2>Overview</h2>

This API was developed to streamline task management for the executive team of the UoA Swedish Club, ensuring the smooth operation of the organisation.
It provides a structured way to create, retrieve, and manage tasks, with features such as task categorisation and due date filtering, helping execs stay organised and on top of their responsibilities.

Built with Go, the API follows RESTful principles and is deployed for real-world use, enabling efficient and scalable task management.
It serves as the backend for our <a>task management dashboard,</a> making it easier to coordinate and oversee club activities.

<h2>Endpoints</h2>


<table>
  <thead>
    <tr>
      <th>Method</th>
      <th>Endpoint</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>POST</td>
      <td>/task/</td>
      <td>Create a new task</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/task/</td>
      <td>Retrieve all tasks</td>
    </tr>
    <tr>
      <td>DELETE</td>
      <td>/task/</td>
      <td>Delete all tasks</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/task/{id}/</td>
      <td>Retrieve a task by ID</td>
    </tr>
    <tr>
      <td>DELETE</td>
      <td>/task/{id}/</td>
      <td>Delete a task by ID</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/tag/{tag}/</td>
      <td>Retrieve tasks by tag</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/due/{year}/{month}/{day}/</td>
      <td>Retrieve tasks due on a specific date</td>
    </tr>
  </tbody>
</table>

<h2>Example Usage</h2>

<h3>Creating a Task</h3>

**Request:**  
To create a new task, send a `POST` request to `/task/` with the following JSON payload:

```json
{
  "text": "Organise Catering for Fika from Scandibunz",
  "tags": ["event", "fika", "scandibunz"], 
  "due": "2025-06-25T14:00:00Z"
}

```


<h2>Features</h2>

- Create, retrieve, delete, and filter tasks
- Hosted and deployed for real-world use, powering our club management
- RESTful API design following best practices  
- CORS-enabled for frontend integration  


<h2>Tech Stack</h2>

- **Go** – Core backend development  
- **HTTP ServeMux** – Request routing  
- **Render.com** – Deployment  
- **Environment Variables** – Secure configuration  




