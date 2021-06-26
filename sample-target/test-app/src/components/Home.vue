<template>
    <div>
        <button name="btn_add_to" onmouseenter="this" onmouseleave="this">Checkout</button>
        <button name="btn_order">Add to basket</button>
        <input type="submit"  value="submit" name="btn_buy" id="">
        <button @click="clicked">Click Me!</button>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    name: 'Home',
    methods: {
        clicked() {
            axios.get("http://localhost:8080/login").then(resp => {console.log(resp)}).catch(err => {console.log(err)});
        }
    }
}
</script>