<template>
<div class="sidebar" :class="{'active': active}">
    <div class="logo_content">
      <div class="logo">
        <div class="logo_name">datalab.dev</div>
      </div>
      <i class='bx bx-menu' id="btn" @click="ellapse()"></i>
    </div>
    <ul class="nav_list">
      <li @click="setActive('view_queries')">
          <i class='bx bx-grid-alt' ></i>
          <span class="links_name">Dashboard</span>
      </li>
      <li @click="setActive('view_app')">
          <i class='bx bx-package' ></i>
          <span class="links_name">Apps</span>
      </li>
      <li @click="setActive('view_account')">
          <i class='bx bx-user' ></i>
          <span class="links_name">User</span>
      </li>
      <li @click="setActive('view_notify')">
          <span v-if="notifications > 0" class="bubble"></span>
          <i class='bx bx-chat' ></i>
          <span class="links_name">Notifications</span>
      </li>
      <li class="no-bg" @click="setActive('view_logout')">
          <i class="icon hover">👋</i>
          <span class="links_name">Logout</span>
      </li>
      <li class="no-mode justify-center">
        <div class="">
            <input id="toggle" class="toggle" type="checkbox" @change="setMode($event)">
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
import MenuItem from '@/components/side_menu/MenuItem.vue';
import jwt_decode from "jwt-decode";

export default {
  name: 'SideMenu',
  data() {
    return {
      mode: "light",
      active: true,
      loggedInUser: {},
    };
  },
  components: {
      MenuItem,
  },
  created() {
    this.loggedInUser = jwt_decode(localStorage.getItem("token"))
    this.mode = localStorage.getItem("theme");
    if (this.mode === undefined || this.mode === null) {
      localStorage.setItem("theme", "light");
      this.mode = "light";
    }
    this.setMode();
  },
  methods: {
    ellapse() {
      this.active = !this.active;
    },
      setActive(value) {
          this.$store.commit('setActiveTab', value);
          this.$emit('setActive', value);
      },
      setMode() {
        var root = document.documentElement
        
        if (this.mode === "light") {
          localStorage.setItem("theme", this.mode)
          this.mode = "dark"
          root.style.setProperty("--main-bg", "#cccccc1B");
          root.style.setProperty("--sub-bg", "#fff");
          root.style.setProperty("--tab-bg", "#fff");
          root.style.setProperty("--btn-font-hover", "#fff");
          root.style.setProperty("--tab-font-selected", "#fff");
          root.style.setProperty("--h-color", "#666666");
          root.style.setProperty("--txt-small", "#666666AA");
          root.style.setProperty("--main-color", "#666666AA");
          root.style.setProperty("--font-blue", "#00000075");
          root.style.setProperty("--font-green", "#00000075");
          root.style.setProperty("--font-yellow", "#00000075");
        } else {
          localStorage.setItem("theme", this.mode)
          this.mode = "light"
          root.style.setProperty("--main-bg", "#1E1E1E");
          root.style.setProperty("--sub-bg", "#2F2F2F");
          root.style.setProperty("--btn-font-hover", "#121212");
          root.style.setProperty("--tab-font-selected", "#121212");
          root.style.setProperty("--tab-bg", "#666666");
          root.style.setProperty("--h-color", "#ccc");
          root.style.setProperty("--txt-small", "#FFFFFFAA");
          root.style.setProperty("--main-color", "#FFFFFFAA");
          root.style.setProperty("--font-blue", "#58A6FF");
          root.style.setProperty("--font-green", "#10d574");
          root.style.setProperty("--font-yellow", "#f7fd04");
        }
      },
  },
  computed: {
    notifications() {
        return this.$store.state.notifications?.length;
      },
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>


::selection{
  color: #fff;
  background: linear-gradient(0deg, #50e3c2 0%,#10d574 100%);;
}
.sidebar{
  grid-column: 1;
  grid-row: 1 / 4;
  /* position: fixed; */
  top: 0;
  left: 0;
  height: 100%;
  width: 78px;
  background: linear-gradient(0deg, #50e3c2 0%,#10d574 100%);;
  padding: 6px 14px;
  z-index: 99;
  transition: all 0.5s ease;
}
.sidebar.active{
  width: 240px
}
.sidebar .logo_content .logo{
  color: #fff;
  display: flex;
  height: 50px;
  width: 100%;
  align-items: center;
  opacity: 0;
  pointer-events: none;
  transition: all 0.5s ease;
}
.sidebar.active .logo_content .logo{
  opacity: 1;
  pointer-events: none;
}
.logo_content .logo i{
  font-size: 28px;
  margin-right: 5px;
}
.logo_content .logo .logo_name{
  font-size: 20px;
  font-weight: 400;
}
.sidebar #btn{
  position: absolute;
  color: #fff;
  top: 6px;
  left: 38px;
  font-size: 22px;
  height: 50px;
  width: 50px;
  text-align: center;
  line-height: 50px;
  transform: translateX(-50%);
  cursor: pointer;
}
.sidebar.active #btn{
  left: 200px;
}
.sidebar ul{
  margin-top: 20px;
}
.sidebar ul li{
  position: relative;
  height: 50px;
  width: 100%;
  margin: 0 5px;
  list-style: none;
  line-height: 50px;
  margin: 5px 0;
}
.sidebar ul li .tooltip{
  position: absolute;
  left: 125px;
  top: 0;
  transform: translate(-50% , -50%);
  border-radius: 6px;
  height: 35px;
  width: 120px;
  background: #fff;
  line-height: 35px;
  text-align: center;
  box-shadow: 0 5px 10px rgba(0, 0, 0, 0.2);
  transition: 0s;
  opacity: 0;
  pointer-events: none;
  display: block;
}
.sidebar.active ul li .tooltip{
  display: none;
}
.sidebar ul li:hover .tooltip{
  transition: all 0.5s ease;
  opacity: 1;
  top: 50%
}

.sidebar ul li{
  color: #fff;
  display: flex;
  align-items: center;
  text-decoration: none;
  border-radius: 12px;
  white-space: nowrap;
  transition: all 0.4s ease;
}

.no-mode:hover {
  background: none !important;
}
.no-mode div {
  position: absolute;
  left: 15px;
  bottom: 30px;
  top: 0;
}
.no-bg {
  position: absolute;
  bottom: 0px;
}
.no-bg:hover {
  background: none !important;
}
.sidebar ul li:hover{
  color: #11101d;
  background: #fff;
  cursor: pointer;
}
.sidebar ul li i{
  font-size: 18px;
  font-weight: 400;
  height: 50px;
  min-width: 50px;
  border-radius: 12px;
  line-height: 50px;
  text-align: center;
}
.sidebar .links_name{
  font-size: 18px;
  font-weight: 400;
  opacity: 0;
  pointer-events: none;
  transition: all 0.3s ease;
}
.sidebar.active .links_name{
  transition: 0s;
  opacity: 1;
  pointer-events: auto
}

.sidebar.active #log_out{
  position: relative;
  top: 75%;
  left: 75%;
  background: none;
  font-size: 25px;
  /* width: 50px; */
}
.sidebar #log_out{
  position: relative;
  top: 75%;
  background: none;
  font-size: 25px;
  /* width: 50px; */
}
.home_content{
  position: absolute;
  height: 100%;
  width: calc(100% - 78px);
  left: 78px;
  background: #E4E9F7;
  box-shadow: 0 5px 10px rgba(0,0,0,0.2px);
  transition: all 0.5s ease;
}
.sidebar.active ~ .home_content{
  z-index: 100;
}
.home_content .text{
  font-size: 25px;
  font-weight: 500;
  color: #10d574;
  margin: 12px;
}
.sidebar.active ~ .home_content{
  width: calc(100% - 240px);
  left: 240px;
}


.toggle {
  --size: 1.4rem;
  -webkit-appearance: none;
     -moz-appearance: none;
          appearance: none;
  outline: none;
  cursor: pointer;
  width: var(--size);
  height: var(--size);
  box-shadow: inset calc(var(--size) * 0.33) calc(var(--size) * -0.25) 0;
  border-radius: 999px;
  color:  #03045e;
  transition: all 500ms;
}
.toggle:checked {
  --ray-size: calc(var(--size) * -0.4);
  --offset-orthogonal: calc(var(--size) * 0.65);
  --offset-diagonal: calc(var(--size) * 0.45);
  transform: scale(0.75);
  color: #ffaa00;
  box-shadow: inset 0 0 0 var(--size), calc(var(--offset-orthogonal) * -1) 0 0 var(--ray-size), var(--offset-orthogonal) 0 0 var(--ray-size), 0 calc(var(--offset-orthogonal) * -1) 0 var(--ray-size), 0 var(--offset-orthogonal) 0 var(--ray-size), calc(var(--offset-diagonal) * -1) calc(var(--offset-diagonal) * -1) 0 var(--ray-size), var(--offset-diagonal) var(--offset-diagonal) 0 var(--ray-size), calc(var(--offset-diagonal) * -1) var(--offset-diagonal) 0 var(--ray-size), var(--offset-diagonal) calc(var(--offset-diagonal) * -1) 0 var(--ray-size);
}
.toggle {
  z-index: 1;
}
.toggle:checked ~ .background {
  --bg: white;
}

.bubble {
  width: 15px;
  height: 15px;
  border-radius: 50%;
  background: #5465ff;
  position: absolute;
  right: -5px;
  top: -5px;
}

</style>
