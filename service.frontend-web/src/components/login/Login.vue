<template>
  <div class="login-page">
    <div class="info-box">
      <h1 class="title-text">DataLab</h1>
      <div>Mine Your Business</div>
    </div>
    <div v-if="mode === 'signin'" class="signinup">
      <div class="d-flex justify-end">
        <div class="prime-text" @click="switchMode('signup')">Sing Up</div>
      </div>
      <div class="d-flex justify-center">
        <div class="login width">
          <div class="creds">
            <div class="form-row">
              <div class="form-group col">
                <input type="text" class="form-control" placeholder="@username" v-model="input.username"/>
              </div>
            </div>
            <div class="form-row">
              <div class="form-group col">
                <input type="password" class="form-control" placeholder="password" v-model="input.password"/>
              </div>
            </div>
            <div class="submit-btn">
              <button @click="login"><span>Sign In</span></button>
            </div>
          </div>
      </div>
      </div>
    </div>
    <div v-if="mode === 'signup'" class="signinup">
      <div class="d-flex justify-end">
        <div class="prime-text" @click="switchMode('signin')">Sing In</div>
      </div>
      <div class="d-flex justify-center">
        <div class="login">
          <div class="creds">
            <!-- <div class="inputs"> -->
              <div class="text-capture big">About You</div>
              <div class="form-row">
                <div class="form-group col">
                  <input type="text" class="form-control" name="first_name" for="first_name" id="first_name" placeholder="First Name" v-model="signup.first_name"/>
                </div>
                <div class="form-group col">
                  <input type="text" class="form-control" name="last_name" for="last_name" id="last_name" placeholder="Last Name" v-model="signup.last_name"/>
                </div>
              </div>
              <div class="form-row">
                <div class="form-group col">
                  <input type="text" class="form-control" name="username" for="username" id="username" placeholder="@username" v-model="signup.username"/>
                </div>
              </div>
              <hr>
              <div class="text-capture big mt-2">Your Company</div>
              <div class="form-row">
                <div class="form-group col">
                  <input type="text" class="form-control" name="orgn_domain" for="orgn_domain" id="orgn_domain" placeholder="Organization domain" v-model="signup.orgn_domain"/>
                </div>
                <div class="form-group col">
                  <input type="text" class="form-control" name="orgn_position" for="orgn_position" id="orgn_position" placeholder="Your Position" v-model="signup.orgn_position"/>
                </div>
              
              </div>
              <hr>
              <div class="text-capture big mt-2">Security</div>
              <div class="form-row">
                <div class="form-group col">
                  <input type="password" class="form-control" name="password" for="password" id="password" placeholder="Password" v-model="signup.password"/>
                </div>
                <div class="form-group col">
                  <input type="password" class="form-control" name="passwordConfirm" for="passwordConfirm" id="passwordConfirm" placeholder="Confirm Password" v-model="signup.passwordConfirm"/>
                </div>
              </div>
            <!-- </div> -->
            <div class="submit-btn">
              <button @click="register"><span>Sign Up</span></button>
            </div>
        </div>
      </div>
      </div>
    </div>
  </div>
</template>

<script>

import axios from 'axios';

export default {
  name: 'Login',
  data() {
    return {
      mode: "signin",
      input: { username: '', password: '' },
      errors: [], 
      hasRegistered: false, 
      signup: { 
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
    login() {
      axios.post('http://192.168.0.177:8080/api/v1/user/login', {
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
    register() {
      if ((this.input.password.length <= 0 || this.input.passwordConfirm.length <= 0) && (this.input.password != this.input.passwordConfirm)) {
        return
      }
      let payload = {
        username: this.signup.username,
        password: this.signup.password,
        orgn_domain: this.signup.orgn_domain,
        first_name: this.signup.first_name,
        last_name: this.signup.last_name,
        orgn_position: this.signup.orgn_position,
      };
      axios.post('http://192.168.0.177:8080/api/v1/user/register', payload).then((resp) => {
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
    switchMode(value) {
      this.mode = value;
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/* h1, h2{
  font-size: 45px;
  text-align: center;
  font-weight: 100;
  background: var(--gradient-green);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
} */

.login-page {
  height: 100%;
  width: 100%;
  display: grid;
  grid-template-columns: 55% auto;
}


.info-box {
  height: 100%;
  width: 100%;
  background: linear-gradient(310deg, #50e3c2 0%,#10d574 100%);
  padding: 45px 25px;
}
.info-box h1 {
  font-size: 6rem;
  color: white;
  display: flex;
  justify-content: flex-start;
}
.info-box div {
  font-size: 4em;
  color: white;
  opacity: 0.8;
  display: flex;
  justify-content: flex-start;
}

.signinup {
  display: grid;
  padding: 15px 25px;
}

.prime-text {
  font-size: 1rem;
  font-weight: bold;
  color: var(--h-color);
  cursor: pointer;
}
.login {
  padding: 25px 25px 0px 25px;
  border-radius: 8px;
  box-shadow: 0 0 11px 0px var(--btn-bg);
  height: max-content;
  display: grid;
  row-gap: 25px;
  align-content: center;
  justify-content: center;
  backdrop-filter: blur(4px); 
}
.login.width {
  width: 300px;
  max-width: 400px;
}

@media only screen and (max-width: 600px) {
  .login-page {
    grid-template-columns: none !important;
    grid-template-rows: 0.6fr auto;
  }
  .info-box {
    grid-row: 2;
  }
  .info-box h1 {
  font-size: 4rem;
  }
  .info-box div {
    font-size: 2em;
  }
}

.creds {
  display: grid;
  justify-content: center;
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
  width: 160px;
  height: 40px;
  font-size: 18px;
  font-weight: 200;
  border: none;

}

.login-footer {
  display: flex;
  justify-content: space-evenly;
  width: 255px;
}

.login-footer .ankor-text {
  font-size: 11px;
  opacity: 0.6;
  color: var(--txt-small);
  margin: 0px 5px;
}

.ankor-text:hover {
  cursor: pointer;
  color: var(--txt-small);
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
