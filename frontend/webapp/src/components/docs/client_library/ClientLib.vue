<template>
    <div class="doc-frame">
        <vue3-markdown-it :source='markdown'/>
    </div>
</template>

<script>
import axios from 'axios';
import 'highlight.js/styles/monokai.css';

export default {
    name: "ClientLib",
    data() {
        return {
            source: "# One more moment..."
        }
    },
    computed: {
        markdown() {
            return this.source;
        },
    },
    created() {
        axios.get("http://192.168.0.177:8000/markdown/client_lib.md", {
            headers: {
                // "Cache-Control": "no-cache,max-age=0"
            }
        }).then(res => {
            this.source = res.data;
            console.log(this.source);
        }).catch(err => {console.log(err)});
    },
}
</script>

<style scoped>
.doc-frame {
    width: 800px;
    height: 800px;
    max-height: 800px;
    overflow-y: scroll;
}
</style>