<template>
  <div class="tile is-parent">
    <div class="tile is-child">
      <div class="nicebox" v-bind:style="boxColor">
        <medium-editor :text=foodData.Name class="title is-5 pad":options="editorOptions" v-on:edit="editName"/>

        <span class="subtitle is-7 rgt pad">
          QTY: {{foodData.Count}}
        </span>
        <p class="description-top">
          Location: Pad {{foodData.PadNum}}
        </p>
        <p class="description">
          Expires: <medium-editor :text=foodData.ExpirationDate.substr(0,10) :options="editorOptions" v-on:edit="editExpir"/>
        </p>
        <p class="description">
          Weight Sensor: {{foodData.Position}}
        </p>
        <p class="description">
          Time on Shelf: {{Number((((Date.now() - foodData.Created)/1000)/60)/60).toFixed(1)}} Hours
        </p>
        <div class="boxfooter">
          Quantity: <button class="button" @click="foodData.Count++;promptUpdate=true;">+</button><button @click="foodData.Count--;promptUpdate=true;" class="button">-</button>
        </div>
        <div class="boxfooter">
          <button @click="remove" class="button" title="Caution: Permanent" v-tippy>Remove</button>
          <button v-if="promptUpdate" class="button is-danger update" @click="updateData">Save</button>
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
      boxColor: {
        backgroundColor: 'rgba(14, 105, 161, 0.03)',
      },
      editorOptions: {
        disableReturn: true,
        extraTime: "",
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

  mounted: function () {
    this.$nextTick(function () {
      var date = new Date()

      if(Date.parse(this.foodData.ExpirationDate) < date.getTime()){
        this.boxColor.backgroundColor="#e74c3c22";
      }
      if(Date.parse(this.foodData.ExpirationDate) < (date.getTime() + 86400000) && Date.parse(this.foodData.ExpirationDate) > (date.getTime())){
        this.boxColor.backgroundColor="#f1c40f22";
      }
    })
  },
  updated: function () {

  },
  methods: {
    clean: function(){
    
      if(this.extraTime == "" || this.extraTime == undefined){
        this.extraTime = this.foodData.ExpirationDate.substr(10)

      }
      console.log(this.extraTime)
      this.foodData.ExpirationDate = this.foodData.ExpirationDate.substr(0,10)

    },
    remove: function () {
      fetch('/remove/' + this.foodData.ID, {
        method: 'post'
      }).then(function (data) {
        return data.text()
      }).then(function (data) {
        console.log(data)
      })
      this.$emit('reload')
    },
    updateData: function () {
      if(this.extraTime == "" || this.extraTime == undefined){
        this.extraTime = this.foodData.ExpirationDate.substr(10)

      }
      let el = this
      this.foodData.ExpirationDate += this.extraTime
      console.log(el.foodData)
      fetch('/update', {
        method: 'post',
        body: JSON.stringify(el.foodData)
      }).then(function (data) {
        return data.text()
      }).then(function (data) {
        console.log(data)
        el.promptUpdate = false
        el.$emit('reload')
      })
    },
    editName: function (operation) {
      if(this.foodData.Name != operation.api.origElements.innerHTML){
        this.foodData.Name = operation.api.origElements.innerHTML
        this.promptUpdate = true
      }
    },
    editExpir: function(operation){
      this.clean()
      console.log(operation.api.origElements.innerHTML)
      if(this.foodData.ExpirationDate.substr(0,10) != operation.api.origElements.innerHTML){
        this.foodData.ExpirationDate = operation.api.origElements.innerHTML
        this.promptUpdate = true
      }
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

  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  transition: background-color .15s;
}
.nicebox:hover{
  font-family: "Raleway" sans-serif;
  border-radius: 4px;
  border:1px solid #E0E0EA;

  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
}
</style>
