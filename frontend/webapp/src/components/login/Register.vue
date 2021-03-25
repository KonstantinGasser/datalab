<template>
  <div class="main-container">
    <div>
      <h1 v-if="!hasRegistered">DataLabs - Mine Your Business</h1>
      <h1 v-if="hasRegistered">Welcome to Datalabs</h1>
    </div>
    <!-- <h1>DataLabs Mine Your business </h1> -->
    <div class="login-grid">
    <div class="context">
      <div class="login-form">
          <div class="creds">
              <div class="inputs">
                <div class="text-capture big">About You</div>
                <div class="field-set">
                  <input type="text" name="first_name" for="first_name" id="first_name" placeholder="First Name" v-model="input.first_name"/>
                  <input type="text" name="last_name" for="last_name" id="last_name" placeholder="Last Name" v-model="input.last_name"/>
                </div>
                <div class="field-set">
                  <input type="text" name="username" for="username" id="username" placeholder="@username" v-model="input.username"/>
                </div>
                <hr>
                <div class="text-capture big mt-2">Your Company</div>
                <div class="field-set">
                  <input type="text" name="orgn_domain" for="orgn_domain" id="orgn_domain" placeholder="Organization domain" v-model="input.orgn_domain"/>
                  <input type="text" name="orgn_position" for="orgn_position" id="orgn_position" placeholder="Your Position" v-model="input.orgn_position"/>
                  <!-- <span class="tooltiptext">this is used to map colleagues to your project<br>(example com.datalabs)</span> -->
                </div>
                <hr>
                <div class="text-capture big mt-2">Add Protection</div>
                <div class="field-set">
                  <input type="password" name="password" for="password" id="password" placeholder="Password" v-model="input.password"/>
                  <input type="password" name="passwordConfirm" for="passwordConfirm" id="passwordConfirm" placeholder="Confirm Password" v-model="input.passwordConfirm"/>
                </div>
              </div>
              <div class="submit-btn">
                <button @click="register"><span>Register for free</span></button>
              </div>
          </div>
      </div>
      <div class="wave">
        <img src="../../assets/wave.svg" alt="" />
      </div>
    </div>
    <div class="footer">
      <div class="ankor-text" @click="switchMode">Sign In</div>
    </div>
  </div>
  </div>
</template>

<script>

import axios from 'axios';

export default {
  name: 'Register',
  data() {
    return { 
      errors: [], 
      hasRegistered: false, 
      input: { 
        username: null,
        first_name: null,
        last_name: null,
        password: null, 
        passwordConfirm: null, 
        orgn_domain: null,
        orgn_position: null,
      },
    };
  },
  methods: {
    register() {
      if ((this.input.password.length <= 0 || this.input.passwordConfirm.length <= 0) && (this.input.password != this.input.passwordConfirm)) {
        return
      }
      let payload = {
        username: this.input.username,
        password: this.input.password,
        orgn_domain: this.input.orgn_domain,
        first_name: this.input.first_name,
        last_name: this.input.last_name,
        orgn_position: this.input.orgn_position,
      };
      axios.post('http://localhost:8080/api/v1/user/register', payload).then((resp) => {
        if (resp.status === 200) {
          this.hasRegistered = !this.hasRegistered;
          this.$toast.success("Your account has been created");
        }
      }).catch(error => {
        if (error.response.data.status >= 500) {
          this.$toast.error("I am sorry ~ something failed here");
          return
        }
        let msg = (!error.response.data.status || error.response.data.status != 400) ? "Whoops something went wrong..." : "mhm looks like this user already exists..try another @username";
        this.$toast.warning(msg);
      });
    },
    switchMode() {
      this.$router.replace({ name: 'login' });
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
  background: linear-gradient(135deg, #50e3c2 0%,#10d574 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}
hr {
  margin: 3px 0;
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
  /* background: #0D1116; */
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
  background: #1E1E1E;
  border-radius: 20px;
  box-shadow: 0 0 16px 10px rgb(0 0 0 / 10%);
}

.login-form {
  display: grid;
  justify-self: baseline;
  grid-template-columns: 500px;
  align-self: center;
  padding-left: 50px;
}

.inputs .field, .field-set {
  margin: 10px 0;
}

.creds {
  display: grid;
  justify-content: center;
}

.inputs .field, .field-set input {
  height: 35px;
  font-size: 18px;
  border-radius: 5px;
  padding: 5px 15px;
  border: 1px solid rgba(0, 0, 0, 0.2);
  width: 200px;
}

.inputs .field-set input {
  margin: 0 3px;
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
  color: #10d574;
  font-size: 18px;
  font-weight: 200;
  transition: all 0.3s ease-in-out;
}
.submit-btn button:hover {
  cursor: pointer;
  background: linear-gradient(135deg, #50e3c2 0%,#10d574 100%);
  color: #121212;
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
  color: #ccc;
  margin: 0px 5px;
}

.ankor-text:hover {
  cursor: pointer;
  color: #fff;
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

.tooltip {
  position: relative;
}

.tooltip .tooltiptext {
  visibility: hidden;
  width: 200px;
  bottom: -35%;
  left: 101%;
  margin-left: 0px; 
  background-color: black;
  opacity: 0.5;
  color: #fff;
  text-align: center;
  border-radius: 6px;
  padding: 5px 0;

  /* Position the tooltip */
  position: absolute;
  z-index: 1;
}

.tooltip:hover .tooltiptext {
  visibility: visible;
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
