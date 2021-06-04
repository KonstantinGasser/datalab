<template>
  <div class="login-page">
    <div class="info-box">
      
    </div>
    <div class="signinup">
      <div class="d-flex justify-end">
        <div class="prime-text" @click="switchMode">Sing In</div>
      </div>
      <div class="d-flex justify-center">
        <div class="login">
          <div class="creds">
            <!-- <div class="inputs"> -->
              <div class="text-capture big">About You</div>
              <div class="form-row">
                <div class="form-group col">
                  <input type="text" class="form-control" name="first_name" for="first_name" id="first_name" placeholder="First Name" v-model="input.first_name"/>
                </div>
                <div class="form-group col">
                  <input type="text" class="form-control" name="last_name" for="last_name" id="last_name" placeholder="Last Name" v-model="input.last_name"/>
                </div>
              </div>
              <div class="form-row">
                <div class="form-group col">
                  <input type="text" class="form-control" name="username" for="username" id="username" placeholder="@username" v-model="input.username"/>
                </div>
              </div>
              <hr>
              <div class="text-capture big mt-2">Your Company</div>
              <div class="form-row">
                <div class="form-group col">
                  <input type="text" class="form-control" name="orgn_domain" for="orgn_domain" id="orgn_domain" placeholder="Organization domain" v-model="input.orgn_domain"/>
                </div>
                <div class="form-group col">
                  <input type="text" class="form-control" name="orgn_position" for="orgn_position" id="orgn_position" placeholder="Your Position" v-model="input.orgn_position"/>
                </div>
              
              </div>
              <hr>
              <div class="text-capture big mt-2">Security</div>
              <div class="form-row">
                <div class="form-group col">
                  <input type="password" class="form-control" name="password" for="password" id="password" placeholder="Password" v-model="input.password"/>
                </div>
                <div class="form-group col">
                  <input type="password" class="form-control" name="passwordConfirm" for="passwordConfirm" id="passwordConfirm" placeholder="Confirm Password" v-model="input.passwordConfirm"/>
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
        organization: this.input.orgn_domain,
        firstname: this.input.first_name,
        lastname: this.input.last_name,
        position: this.input.orgn_position,
      };
      axios.post('http://192.168.178.103:8080/api/v1/user/register', payload).then((resp) => {
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
  background: var(--gradient-green);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.login-page {
  height: 100%;
  width: 100%;
  display: grid;
  grid-template-columns: 65% auto;
}

.info-box {
  height: 100%;
  width: 100%;
  background: linear-gradient(310deg, #50e3c2 0%,#10d574 100%);
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
  box-shadow: 0 0 11px 0px rgb(0 0 0 / 10%);
  height: max-content;
  display: grid;
  row-gap: 25px;
  align-content: center;
  justify-content: center;   
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
  width: 180px;
  height: 40px;
  font-size: 18px;
  font-weight: 200;
  border: none;

}

.login-footer {
  display: flex;
  justify-content: flex-start;
}

.login-footer .ankor-text {
  font-size: 14px;
  opacity: 0.6;
  color: var(--txt-small);
  margin: 0px 5px;
}

.ankor-text:hover {
  cursor: pointer;
  color: var(--txt-small);
}
</style>
