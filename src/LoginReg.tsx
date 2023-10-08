import {Accessor, Setter, Show} from 'solid-js';
import './styles/loginReg.css';

interface LoginSignal {
  loggedIn: Accessor<boolean>;
  setLoggedIn: Setter<boolean>;
}

function LoginReg(props: LoginSignal) {

  const logInInterface = (
    <form class="login-container">
      <h1>HackUTA23</h1>
      <h3 class="field-description">Email</h3>
      <input type="text" placeholder="example@email.com"/>
      <h3 class="field-description">Password</h3>
      <input type="text" placeholder="password"/>
      <section class="button-container">
        <button type="submit" id="login-button">Login</button>
        <button type="submit" id="signup-button">Sign Up</button>
      </section>
    </form>
    );

  return (
    <>
      <Show when={props.loggedIn()}
            fallback={logInInterface}>
        <p>You've logged in successfully</p>
      </Show>
    </>
  );
}

export default LoginReg;