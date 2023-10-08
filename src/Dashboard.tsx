import { For } from 'solid-js';
import './styles/styles.css';
import './styles/dashboard.css'
import { A } from '@solidjs/router';
import { createSignal, createEffect } from 'solid-js';

function Dashboard () {
  const [data, setData] = createSignal(null);

  const fetchData = async function(){    
    const response = await fetch("http://localhost:8080/getall");
    const jsonData = await response.json();
    setData(jsonData);
    console.log(data());
  };

  createEffect(() =>
  {
    const intervalID = setInterval(() =>
    {
      fetchData();
    }, 5000);

    return () => {
      clearInterval(intervalID);
    }
  })

  return(
    <>
      <h1>Alert Log</h1>
      <section class="alert-log-container">
        <For each={data()}>{(dataAlert, i) =>
          <p class="alert-information">
            {dataAlert.Alert + " : " + dataAlert.Time}
          </p>
        }</For>
      </section>
      <A href="/" style="margin-top: 1em">Sign Out</A>
    </>
  );
}

export default Dashboard;