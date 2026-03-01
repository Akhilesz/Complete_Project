import logo from './logo.svg';
import react from 'react';
import './App.css';

function App() {
  const [todos , setTodos] = react.useState([]);
  const [newTodo , setNewTodo] = react.useState('');


  addNewTodo = () => {
      
  }
  return (
    <div className="App">
      <h1>
        Todo List
      </h1>
    </div>
  );
}

export default App;
