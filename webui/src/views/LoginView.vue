<script>
export default{
  data(){
    return{
      username: "",
      errormsg: ""
    }
  },
  methods:{
    async trovaUtente(){
        try{
            let response = await this.$axios.post("/session",{username: this.username});
            var profilo = response.data;
            localStorage.setItem('identifier', profilo.UserID);
            localStorage.setItem('username', profilo.Username);
            this.$router.replace('/session');
        }
        catch(e)
        {
            this.errormsg = e.toString();
            alert(this.errormsg);
            console.log(this.errormsg);
        }
    }
  }
}
</script>

<template>
  <title>Schermata Login</title>

  <body>
    <img class = sfondo src= 'https://www.ifolor.ch/content/dam/ifolor/inspire-gallery/FFL/photoshop-ein-panorama-aus-einzelaufnahmen-erstellen/Panorama-580.jpg.transform/q60/image.jpg?Panorama-580.jpg'>
    <h1 class = titolo >WasaPhoto</h1>
    <input class = login type ="text" placeholder="Inserisci il tuo username" v-model="username">
    <button class= accesso :disabled= "username.length > 16 || username.length < 3"
      @click = "trovaUtente">Accedi</button>
  </body>
  
</template>

<style>
  body{
    width:100%;
    height: 100%;
    border: none;
    position:absolute;
    margin-left: auto;
    margin-top: auto;
  }

  .sfondo{
    position: absolute;
    width:100%;
    height:100%;
  }

  .titolo{
    position: absolute;
    transform: translate(-50%,-50%);

    left: 50%;
    top: 30%;
    font-size: 800%;

    -webkit-text-stroke-width: 0.5px;
    -webkit-text-stroke-color: black;

    font-family:'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
    font-style:italic;
    color: white;
  }

  .accesso{
    position:absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%,-50%);

    width:200px;
    height:30px;
    
    cursor: pointer;
    text-align:center;
    border-radius: 5px;
    font-family:'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
    background-color: rgba(255, 255, 255, 0.9);
  }

  .login{
    position: absolute;
    top: 43%;
    left: 50%;
    transform: translate(-50%,-50%);

    width:300px;
    height:30px;
    font-size: 25px;

    cursor: pointer;
    align-self:right;
    border-radius: 5px;
    background-color: rgba(0, 0, 0, 0.5);
    
    color: white;
    font-family:'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
    text-align: center;
  }


</style>