<template>
  <div class="main-container">
    <div>
      <h1>Datalab - Mine Your Business</h1>
    </div>
    <!-- <h1>datalab Mine Your business </h1> -->
    <div class="login-grid">
    <div class="context">
      <div class="login-form">
          <div class="creds">
            <div class="inputs">
              <div class="field">
                <input type="text" placeholder="@username" v-model="input.username"/>
              </div>
              <div class="field">
                <input type="password" placeholder="password" v-model="input.password"/>
              </div>
            </div>
            <div class="submit-btn">
              <button @click="login"><span>SignIn</span></button>
            </div>
          </div>
      </div>
      <div class="wave">
        <img src="../../assets/wave.svg" alt="" />
      </div>
    </div>
    <div class="footer">
      <div class="ankor-text">Forgot Password?</div>
      <div class="ankor-text">|</div>
      <div class="ankor-text" @click="switchMode">Register for free</div>
    </div>
  </div>
  </div>
</template>

<script>

import axios from 'axios';

export default {
  name: 'Login',
  data() {
    return { input: { username: '', password: '' } };
  },
  methods: {
    login() {
      axios.post('http://localhost:8080/api/v1/user/login', {
        username: this.input.username,
        password: this.input.password,
      }).then(resp => {
          if (resp.status === 200) {
            localStorage.setItem('token', resp.data.token);
            this.$router.replace({ name: 'dashboard' });
          }
        }).catch((err) => {
          if (err.response.status === 403) {
            this.$toast.warning("Mhm looks like the username || password is wrong..");
          }
          if (err.response.status === 500) {
            this.$toast.warning("Mhm sorry something did not work ...");
          }
        });
    },
    switchMode() {
      this.$router.replace({ name: 'register' });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2{
  font-size: 45px;
  text-align: center;
  font-weight: 100;
  background: var(--gradient-green);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.main-container {
  height: 100%;
  min-height: 100vh;
  width: 100vw;
  display: grid;
  justify-content: center;
  align-content: center;
  grid-template-rows: min-content min-content;
  grid-row-gap: 25px;
}
.component-fade-enter-active, .component-fade-leave-active {
  transition: opacity .3s ease;
}
.component-fade-enter, .component-fade-leave-to
/* .component-fade-leave-active for <2.1.8 */ {
  opacity: 0;
}
.login-grid {
  /* display: grid;
  grid-template-rows: 100px max-content; */
}
.context {
  display: flex;
  height: 471px;
  background: var(--sub-bg);
  border-radius: 20px;
  box-shadow: 0 0 16px 10px rgb(0 0 0 / 10%);
}

.login-form {
  display: grid;
  justify-self: baseline;
  grid-template-columns: 300px;
  align-self: center;
  padding-left: 50px;
  height: 175px;
}

.inputs .field {
  margin: 10px 0;
}

.creds {
  display: grid;
  justify-content: center;
}

.inputs .field input {
  height: 35px;
  font-size: 18px;
  border-radius: 5px;
  padding: 5px 15px;
  border: 1px solid var(--sub-border);
  width: 200px;
}

.submit-btn {
  margin-top: 10px;
  display: flex;
  justify-content: center;
  height: 50px;
  
}

.submit-btn button {
  width: 150px;
  height: 35px;
  border-radius: 50px;
  background-color: transparent;
  border: none; 
  font-size: 18px;
  font-weight: 200;
  transition: all 0.3s ease-in-out;
}
.submit-btn button:hover {
  cursor: pointer;
  background: var(--gradient-green);
  color: var(--btn-font-hover);
  width: 180px;
  height: 40px;
  font-size: 18px;
  font-weight: 200;
  border: none;

}

.footer {
  margin: 15px 0px;
  display: flex;
  justify-content: flex-start;
}

.footer .ankor-text {
  font-size: 14px;
  opacity: 0.6;
  color: var(--txt-small);
  margin: 0px 5px;
}

.ankor-text:hover {
  cursor: pointer;
  color: var(--txt-small);
}

button {
  outline: none;
}

.wave {
  width: 449px;
}

.wave img {
  object-fit: contain;
  height: 472px;
  border-radius: 20px;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}

input {
  outline: none;
  border: none;
}

input:focus::-webkit-input-placeholder {
  color: transparent;
}
input:focus:-moz-placeholder {
  color: transparent;
}
input:focus::-moz-placeholder {
  color: transparent;
}
input:focus:-ms-input-placeholder {
  color: transparent;
}

textarea:focus::-webkit-input-placeholder {
  color: transparent;
}
textarea:focus:-moz-placeholder {
  color: transparent;
}
textarea:focus::-moz-placeholder {
  color: transparent;
}
textarea:focus:-ms-input-placeholder {
  color: transparent;
}

input::-webkit-input-placeholder {
  color: #999999;
}
input:-moz-placeholder {
  color: #999999;
}
input::-moz-placeholder {
  color: #999999;
}
input:-ms-input-placeholder {
  color: #999999;
}

textarea::-webkit-input-placeholder {
  color: #999999;
}
textarea:-moz-placeholder {
  color: #999999;
}
textarea::-moz-placeholder {
  color: #999999;
}
textarea:-ms-input-placeholder {
  color: #999999;
}
</style>
