<template>
    <div>
        <vue3-markdown-it :source='markdown'/>
    </div>
</template>

<script>
import axios from 'axios'

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
        axios.get("http://localhost:8000/markdown/client_lib.md", {
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