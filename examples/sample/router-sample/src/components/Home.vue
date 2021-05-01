<template>
    <div>
        <button name="checkout" onmouseenter="this" onmouseleave="this">Checkout</button>
        <button name="add-basket">Add to basket</button>
        <input type="submit"  value="submit" name="submit-btn" id="">
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