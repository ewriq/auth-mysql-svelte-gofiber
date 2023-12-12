<script>
  // @ts-ignore
  import { Router, Link, Route } from "svelte-routing";
  import Home from "./pages/Home.svelte";
  import Login from "./pages/Login.svelte";
  import Register from "./pages/Register.svelte";
  // @ts-ignore
  import Dashboard from "./pages/Dahboard.svelte";
  // @ts-ignore
  import Cookies from "js-cookie";
  import axios from "axios";
  

  let pagepath = window.location.pathname;
  let Users = false;
  let user = Cookies.get("token");

  async function GetData() {
    if (user) {
      try {
        const response = await axios.post(
          `http://localhost:3000/api/v1/user/`, {
            token: user
          }
        );

        if (response.data.status === "OK") {
          Users = response.data.data;
          console.log(response.data.data[0].token);
        } else {
          console.error("Kullanıcı hatası ?");
        }
      } catch (err) {
        console.error(err);
      }
    } else {}
  }
  GetData();

  export let url = "";
</script>

<main>
  <Router {url}>
    <nav>  
    <Link to="/register">Register</Link>
    <Link to="/login">Login</Link>
    <Link to="/">Home</Link>
    </nav>

    <div>
      {#if Users}
      <Route path="/dashboard"><Dashboard users={Users} /></Route>
    {/if}
      <Route path="/"><Home /></Route>
      <Route path="/login"><Login /></Route>
      <Route path="/register"><Register /></Route>
    </div>
  </Router>
</main>