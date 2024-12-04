import React from 'react';
import "./App.css"
import {Contrainer} from 'sementic-ui-react'
import ToDoList from './ToDoList'

function App() {
  return (
    <div>
      <Contrainer>
        <ToDoList/>
      </Contrainer>
    </div>
  )
}

export default App

