import {createSignal} from 'solid-js';
import './styles/styles.css';
import './styles/loginReg.css';
import { A } from '@solidjs/router';

function Login() {
  const [email, setEmail] = createSignal("");
  const [password, setPassword] = createSignal("");

  const handleLogin = (e: Event) => {
    e.preventDefault();
    console.log(email() + ' ' + password());
  };

  return (
    <>
      <form class="login-container" onSubmit={handleLogin}>
        <h1>HackUTA23</h1>
        <h3 class="field-description">Email</h3>
        <input type="text" placeholder="example@email.com" value={email()} onInput={(e) => setEmail(e.target.value)}/>
        <h3 class="field-description">Password</h3>
        {<input type="password" placeholder="password" value={password()} onInput={(e) => setPassword(e.target.value)}/>}
        <section class="button-container">
          <button type="submit" id="login-button">Login</button>
          <A href="/signup">Sign Up</A>
        </section>
      </form>
    </>
  );
}

export default Login;