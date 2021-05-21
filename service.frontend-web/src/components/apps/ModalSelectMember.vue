<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <slot name="header"> Add member from your company </slot>
          </div>

          <div class="modal-body">
            <div class="d-flex justify-end align-center"><span class="icon icon-plus big hover green" @click="pushItem()"></span></div>
            <ul class="add_item_list">
                <li class="d-flex justify-between align-center">
                    <div class="w-25">#</div>
                    <div class="w-50">Member</div>
                    <div class="w-25 d-flex justift-center">Action</div>
                </li>
                <hr>
                <li v-for="(item, i) in items" :key="i" class="my-1 d-flex justify-between align-center">
                    <div class="w-25">#{{i+1}}</div>
                    <div class="w-50">
                        <select name="" id="" @change="select(i ,$event)">
                          <option value="Select Member">Select Member</option>
                          <option v-for="m in member_list" :value="m.id" :key="m.id">{{m.id}}</option>
                        </select>
                    </div>
                    <div class="w-25 d-flex justify-center"><span v-if="i !== 0" class="icon icon-delete green hover" @click="dropItem(item.uuid)"></span></div>
                </li>
            </ul>
          </div>

          <div class="modal-footer d-flex justify-end">
                {{items}}
                <button class="btn-standard" @click="$emit('close')">Cancel</button>
                <button v-if="list_valid" class="btn-standard"><span class="icon icon-plus"></span> Add</button>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>
<script>
    export default {
    name: "selectMember",
    props: ["member_list"],
    data() {
        return {
            items: [{uuid: ''}],
            list_valid: false,
        };
    },
    methods: {
      pushItem() {
        this.items.push({uuid: ''});
      },
      dropItem(uuid) {
        const tmp = []
        for (let i = 0; i < this.items.length; i++) {
          if (this.items[i].uuid !== uuid) { tmp.push(this.items[i]);}
        }
        this.items = tmp;
      },
      select(index, event) { 
        if (event.target.value === 'Select Member') return
        this.items[index].uuid = event.target.value;
        this.$props.member_list = this.$props.member_list.filter(item => {item !== event.target.value});
        this.list_valid = true;
      },
    },
    };
</script>

<style scoped>
</style>
