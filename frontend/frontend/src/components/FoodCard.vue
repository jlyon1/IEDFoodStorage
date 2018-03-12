<template>
  <div class="tile is-parent">
    <div class="tile is-child">
      <div class="nicebox">
        <medium-editor :text=foodData.Name class="title is-5 pad":options="editorOptions" v-on:edit="editName"/>

        <span class="subtitle is-7 rgt pad">
          QTY: {{foodData.Count}}
        </span>
        <p class="description-top">
          Location: Pad {{foodData.PadNum}}
        </p>
        <p class="description">
          Expires: {{foodData.ExpirationDate.substr(0,10)}}
        </p>
        <p class="description">
          Weight Sensor: none
        </p>
        <p class="description">
          Time on Shelf:
        </p>
        <div class="boxfooter">
          Quantity: <button class="button" @click="foodData.Count++;promptUpdate=true;">+</button><button @click="foodData.Count--;promptUpdate=true;" class="button">-</button>
        </div>
        <div class="boxfooter">
          <button @click="remove" class="button" title="Caution: Permanent" v-tippy>Remove</button>
          <button class="button">Move</button>
          <button class="button">Add to List</button>
          <button v-if="promptUpdate" class="button is-danger update" @click="updateData">update</button>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
import editor from 'vue2-medium-editor'

export default {
  name: 'FoodCard',
  props: ['foodData'],
  data () {
    return {
      editorOptions: {
        disableReturn: true,
        toolbar: {
          buttons: []
        }
      },
      promptUpdate: false
    }
  },
  components: {
    'medium-editor': editor
  },
  mounted: function(){
  },
  methods: {
    remove: function(){
      this.$emit('reload')
      fetch('http://127.0.0.1:8081/remove/' + this.foodData.ID, {
        method: 'post'
      }).then(function(data){
        return data.text();
      }).then(function(data){
        console.log(data);
      });
    },
    updateData: function(){
      let el = this;
      console.log(el.foodData)
      fetch('http://127.0.0.1:8081/update', {
        method: 'post',
        body:JSON.stringify(el.foodData)
      }).then(function(data){
        return data.text();
      }).then(function(data){
        console.log(data);
        el.promptUpdate = false;
        el.$emit('reload')

      });
    },
    editName: function(operation){
      this.foodData.Name = operation.api.origElements.innerHTML
      this.promptUpdate = true
    }
  }
}
</script>

<style scoped>
.medium-editor-action {
  display: none;
}
.update{
  transition: display 1s;
}

.small{
  font-size: 20px;
  margin: 5px;
}

.boxfooter{
  padding: 10px;
}
.pad{
  padding-left: 10px;
  padding-right: 10px;
  padding-top: 10px;
  padding-bottom: 0px;

}
.description{
  margin-top: 0;
  margin-left: 0;
  border-top: 0px;
  border-bottom: 1px;
  border-left: 0;
  border-right: 0;
  border-style: solid;
  border-color:  #E0E0EA;
  background-color: white;
  width: 100%;
  position: inherit;
  padding: 15px;
}

.description-top{

  margin-top: 0;
  margin-left: 0;
  border-top: 1px;
  border-bottom: 1px;
  border-left: 0;
  border-right: 0;
  border-style: solid;
  border-color:  #E0E0EA;
  background-color: white;
  width: 100%;
  position: inherit;
  padding: 15px;

}

.rgt{
  float:right;
}
.nicebox{
  font-family: "Raleway" sans-serif;
  border-radius: 4px;
  border:1px solid #E0E0EA;
  background-color: rgba(14, 105, 161, 0.03);
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  transition: background-color .15s;
}
.nicebox:hover{
  font-family: "Raleway" sans-serif;
  border-radius: 4px;
  border:1px solid #E0E0EA;
  background-color: rgba(14, 105, 161, 0.05);
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
}
</style>
