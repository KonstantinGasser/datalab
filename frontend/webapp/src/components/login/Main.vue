<template>
  <div class="main-container">
       <Login @setActive="show" v-if="view === 'login'"/>
       <Register @setActive="show" v-if="view === 'register'"/>
  </div>
</template>

<script>
import Login from './Login.vue';
import Register from './Register.vue';

export default {
  name: 'Main',
  components: {
      Login,
      Register,
  },
  data() {
    return { view: 'login' };
  },
  methods: {
      show(value) {
          this.view = value;
      },
  },
};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h2 {
  font-size: 34px;
  margin: 0;
  padding: 0;
  text-align: center;
  color: #fc682d;
}
.main-container {
  height: 100%;
  min-height: 100vh;
  width: 100vw;
  display: flex;
  justify-content: center;
  align-items: center;
  /* background: rgba(0, 0, 0, 0.1); */
  /* background: linear-gradient(40deg,
            rgba(252,104,45,1) 0%,
            rgba(39,5,85,1) 100%,
            rgba(242,5,87,1) 100%); */
  background: linear-gradient(135deg, #f8f8f8 0%,#efefef 100%);
}
.component-fade-enter-active, .component-fade-leave-active {
  transition: opacity .3s ease;
}
.component-fade-enter, .component-fade-leave-to
/* .component-fade-leave-active for <2.1.8 */ {
  opacity: 0;
}
</style>
