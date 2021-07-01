<template>
    <div style="width: 60%;margin-left:auto;margin-right:auto;">
        <ul style="display:flex;justify-content:space-evenly;">
            <li><router-link to="/basket">Shopping Basket</router-link></li>
        </ul>
        <h1>Sales</h1>
        <div v-for="p in products" :key="p.id" style="border:1px solid;padding:15px;margin:25px 0;">
            <h4>{{p.name}}</h4>
            <p>{{p.desc}}</p>
            <button name="btn_go_to">go to offer</button>
        </div>
        <hr>
        <h1>Trending</h1>
        <div v-for="p in products" :key="p.id" style="border:1px solid;padding:15px;margin:25px 0;">
            <h4>{{p.name}}</h4>
            <p>{{p.desc}}</p>
            <button>view</button>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    name: 'Home',
    data() {
        return {
            products: [
                {id: 1, name: "product#1", desc: "some description"},
                {id: 3, name: "product#3", desc: "some description"},
            ],
        }
    },
}
</script>