<template>
  <main>
    <div class="walpaper">
      <img src="@/assets/img/chartjs-logo.svg" />
      <h3>Budget Manager</h3>
    </div>
    <div class="form">
      <form class="group login active" @submit="submitLogin">
        <input class="username" type="text" placeholder="username" required />
        <input
          class="password"
          type="password"
          placeholder="password"
          required
        />
        <button>
          login <img src="@/assets/img/loader.gif" v-if="login == 1" />
        </button>
        <p class="message">
          Not registered? <a @click="clickMSG">Create an account 👇</a>
        </p>
      </form>
      <form class="group register">
        <input type="text" placeholder="name" required />
        <input type="password" placeholder="password" required />
        <input type="password" placeholder="confirm password" required />
        <button>
          create <img src="@/assets/img/loader.gif" v-if="login == 'pending'" />
        </button>
        <p class="message">
          Already registered? <a @click="clickMSG">Sign In 👆</a>
        </p>
      </form>
    </div>
  </main>
</template>

<script>
import store from "@/store";

export default {
  name: "loginView",
  computer: {
    login() {
      return store.getters["login/get_status_login"];
    },
  },
  methods: {
    clickMSG() {
      document.querySelectorAll(".group").forEach((item) => {
        item.classList.toggle("active");
      });
    },
    submitLogin() {
      let login_form = document.querySelector(".group.login");
      let username = login_form.querySelector(".username").value;
      let password = login_form.querySelector(".password").value;
      store.dispatch("login/request", { username, password });
    },
  },
};
</script>

<style scoped>
main {
  width: 100%;
  height: 100%;
  background-color: #3c3c3c;
  padding: 0px 0px;
  margin: 0px 0px;
  display: flex;
}

main div.walpaper {
  width: 50vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.walpaper h3 {
  font-family: "quicksand";
  color: #ffff;
  font-size: 1.5em;
}
.form {
  width: 50vw;
  height: 100vh;
  overflow-y: hidden;
}

.form form.group.active {
  height: 100vh;
}
.form form.group {
  height: 0vh;
  overflow-y: hidden;
  transition: 0.5s;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
}

.form input {
  font-family: "quicksand bold";
  outline: 0;
  background: #f2f2f2;
  width: 350px;
  border: 0;
  margin: 0 0 40px;
  padding: 15px;
  box-sizing: border-box;
  font-size: 14px;
  border-radius: 50px;
}
.form button {
  font-family: "quicksand bold";
  text-transform: uppercase;
  outline: 0;
  background: #fa6976;
  width: 250px;
  border: 0;
  padding: 15px;
  color: #ffffff;
  font-size: 14px;
  -webkit-transition: all 0.3 ease;
  transition: all 0.3 ease;
  cursor: pointer;
  border-radius: 50px;
}
.form button:hover,
.form button:active,
.form button:focus {
  background: #d55863;
}
.form .message {
  margin: 15px 0 0;
  color: #b3b3b3;
  font-size: 1em;
  font-family: "quicksand bold";
}
.form .message a {
  color: #fa6976;
  text-decoration: none;
  border-bottom: 1px solid #3c3c3c;
  transition: 0.2s;
  cursor: pointer;
}
.form .message a:hover {
  border-bottom: 1px solid #fa6976;
}

.form .group.active button img {
  opacity: 1;
}

.form .group button img {
  opacity: 0;
  width: 40px;
  position: absolute;
  margin-left: 110px;
  top: 57vh;
}

.form .group.register button img {
  top: 63vh;
}
</style>
