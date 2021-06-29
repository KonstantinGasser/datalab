<template>
    <div style="width: 60%;margin-left:auto;margin-right:auto;">
        <h1>Shopping Basket</h1>
        <ul>
            <li v-for="p in products" :key=p.id :name="'product_'+p.id" style="margin:10px 0;padding:5px 5px;border:1px solid;cursor:pointer;display:flex;justify-content:space-evenly;">
                <div><strong>{{p.name}}</strong></div>
                <div> {{p.desc}}</div>
            </li>
        </ul>
        <div>
            <button name="btn_buy">Buy NOW</button>
        </div>
    </div>
</template>

<script>

export default {
    name: 'Basket',
    data() {
        return {
            products: [
                {id: 2, name: "product#2", desc: "some description"},
                {id: 4, name: "product#4", desc: "some description"},
            ],
        }
    },
}
</script>