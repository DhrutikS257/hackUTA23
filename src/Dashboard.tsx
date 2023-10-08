import { For } from 'solid-js';
import './styles/styles.css';
import './styles/dashboard.css'
import { A } from '@solidjs/router';

function Dashboard () {
  let dataAlerts = [["Bruh", "2am"],["Bruh", "2am"],["Bruh", "2am"],["Bruh", "2am"],["Bruh", "2am"]];



  return(
    <>
      <h1>Alert Log</h1>
      <section class="alert-log-container">
        <For each={dataAlerts}>{(dataAlert, i) =>
          <p class="alert-information">{dataAlert[0]}{dataAlert[1]}</p>
        }</For>
      </section>
      <A href="/" style="margin-top: 1em">Sign Out</A>
    </>
  );
}

export default Dashboard;