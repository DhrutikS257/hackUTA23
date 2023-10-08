import { createSignal } from 'solid-js'
import './styles/App.css'
import LoginReg from './LoginReg.tsx'

function App() {
  const [loggedIn, setLoggedIn] = createSignal(false);
 
  return (
    <>
      <LoginReg loggedIn={loggedIn} setLoggedIn={setLoggedIn}/>
    </>
  )
}

export default App
