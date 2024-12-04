import React from 'react';
import "./App.css"
import 'semantic-ui-css/semantic.min.css'; // Import Semantic UI CSS
import { Container } from 'semantic-ui-react'; // Import Container from Semantic UI React
import ToDoList from './ToDoList'

function App() {
  return (
    <div>
      <Container>
        <ToDoList/>
      </Container>
    </div>
  )
}

export default App

