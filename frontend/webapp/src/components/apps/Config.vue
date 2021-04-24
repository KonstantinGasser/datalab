<template>
    <div>
        <h1>Funnel Configuration</h1>
        <div class="view_component">
            <div class="form-row">
                <div class="form-col col">
                    <!-- <small class="info_txt"> -->
                        In order to track which customer acts in which stage of the funnel,
                        you must provide information about the stage and their transitions to the next stage 😬
                    <!-- </small> -->
                </div>
            </div>
            <div class="form-row mt-2">
                <div class="form-col col">
                    <strong>"Stage Name"</strong>: is the name you want the stage to have, it has no effect on the logic 
                    <br>
                    <strong>"Transition"</strong>: this is <strong>important</strong>. Only with the transition we are able to track when
                    a customer jumps into the next stage <small>(rn "Transition" must be the name of the HTML-Element)</small> 
                </div>
            </div>
            <div class="form-row mt-3">
                <div class="form-col col d-flex">
                    <div v-for="f in funnel" :key="f.id" class="d-flex align-center m-1">
                        <div class="funnel">
                            <div class="d-flex justify-end trash-span">
                                <span v-if="f.id === funnel.length - 1" class="icon icon-trash-2 hover" @click="removeStage(f.id)"></span>
                            </div>
                            <div class="d-flex justify-center align-center flex-col">
                                <div class="stage-name">{{f.name}}</div>
                                <div class="stage-transition">{{f.transition}}</div>
                            </div>
                        </div>
                        <div>
                             <span v-if="f.id < funnel.length - 1" class="icon icon-chevron-right super"></span>
                         </div>
                    </div>
                     <div class="funnel add-box d-flex align-center justify-even">
                         <div>
                             <span class="icon icon-chevrons-right super"></span>
                         </div>
                         <div>
                             <div class="">
                                <div class=" col">
                                    <input v-model="stage_name" class="form-control" type="text" name="stage_name" id="stage_name" placeholder="Stage Name">
                                </div>
                            </div>
                            <div class="mt-1">
                                <div class=" col">
                                    <input v-model="stage_transition" class="form-control" type="text" name="stage_tansition" id="stage_tansition" placeholder="Transition">
                                </div>
                            </div>
                         </div>
                         <div>
                             <span class="icon icon-plus hover super" @click="addStage"></span>
                         </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "Configuration",
    data() {
        return {
            stage_name: null,
            stage_transition: null,
            stage_count: 0,
            funnel: [],
        };
    },
    methods: {
        addStage() {
            this.funnel.push({
                id: this.stage_count,
                name: this.stage_name,
                transition: this.stage_transition,
            });
            this.stage_count++
            this.stage_name = null;
            this.stage_transition = null;
        },
        removeStage(id) {
            this.funnel = this.funnel.filter(item => item.id != id);
        },
    },
}
</script>

<style sceped>
.view_component {
    margin-top: 15px;
    padding: 15px;
    border-radius: 8px;
    height: max-width;
}

.funnel {
    padding: 10px;
    width: auto;
    height: 100px;
    min-width: 150px;
    max-width: 200px;
    background: var(--sub-bg);
    border: 1px solid var(--sub-border);
    border-radius: 8px;
    margin: 0 5px;
}

.funnel .stage-name {
    font-size: 18px;
    color: var(--h-color);
}
.funnel .stage-transition {
    font-size: 14px;
    color: var(--txt-small);
}

.add-box {
    opacity: 0.5;
    border: none;
    width: 220px;
    max-width: 220px;
}
.add-box:focus,.add-box:hover {
    opacity: 1;
}


.trash-span {
    height: 14px;
}
</style>