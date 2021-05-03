<template>
    <div class="docs-grid">
        <div class="menu-bar">
            <h2>App Configs</h2>
            <ul>
                <li class="menu-item" @click="jump('one')">Funnel</li>
                <li class="menu-item" @click="jump('two')">Campaign Tracking</li>
                <li class="menu-item" @click="jump('three')">Interesting Button</li>
            </ul>
            <hr>
            <h2>Client Lib</h2>
            <ul>
                <li class="menu-item" @click="jump('one')">Quick Start</li>
                <li class="menu-item" @click="jump('three')">How to Implementation</li>
                <li class="menu-item" @click="jump('three')">Common Gotchas</li>
            </ul>
        </div>
        <div class="menu-content">
            <div id="one">Lorem, ipsum dolor sit amet consectetur adipisicing elit. Fuga officiis, quisquam perspiciatis asperiores sed delectus nulla, tempora assumenda rerum repellendus doloremque. Quam ipsam harum consectetur culpa alias minus error impedit!
            </div>
            <div id="two">Lorem, ipsum dolor sit amet consectetur adipisicing elit. Fuga officiis, quisquam perspiciatis asperiores sed delectus nulla, tempora assumenda rerum repellendus doloremque. Quam ipsam harum consectetur culpa alias minus error impedit!
            </div>
            <div id="three">Lorem, ipsum dolor sit amet consectetur adipisicing elit. Fuga officiis, quisquam perspiciatis asperiores sed delectus nulla, tempora assumenda rerum repellendus doloremque. Quam ipsam harum consectetur culpa alias minus error impedit!
            </div>
        </div>
        <!-- <div v-if="!hasToken" class="view_component">
            <div class="form-row">
                <div class="form-group col">
                    <div class="info-txt">
                        Mhm looks like you have not yet created and application token. You need one of those
                        in order to get the client side working... <a @click="navCreateToken">go - create one</a> 😉
                    </div>
                </div>
            </div>
        </div> -->
    </div>
</template>

<script>
export default {
    name: "DocClientLib",
    data() {
        return {};
    },
    props: ['hasToken'],
    methods: {
        navCreateToken() {
            this.$emit("goCreateToken");
        },
        jump(to) {
            const el = document.getElementById(to);
            el.scrollIntoView(true);
        }
    },
}
</script>

<style scoped>
.docs-grid {
    display: grid;
    grid-template-columns: 25% 1fr;
    column-gap: 0.5em;
    max-height: 70vh;
    height: 70vh
}
.menu-bar {
    grid-column: 1;
    height: min-content;
    min-height: 500px;
    background: var(--sub-bg);
    border: 1px solid var(--sub-border);
    padding: 1em;
    border-radius: 8px;
}
.menu-item {
    font-size: 1rem;
    padding-left: 10px;
    transition: padding 250ms ease;
}
.menu-item:hover {
    cursor: pointer;
    color: var(--btn-bg-hover);
    padding-left: 15px;
}


.menu-content {
    grid-column: 2;
    background: var(--sub-bg);
    border: 1px solid var(--sub-border);
    border-radius: 8px;
    padding: 1em;
    max-height: 100%;
    height: 100%;
    overflow-y: auto;
}

#one {
    height: 500px;
    padding: 2em;
}
#two {
    height: 500px;
    padding: 2em;
}
#three {
    height: 500px;
    padding: 2em;
}

.view_component {
    margin-top: 15px;
    padding: 15px;
    border-radius: 8px;
    height: max-width;
}
</style>