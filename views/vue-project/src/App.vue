<script>
export default {
  data() {
    return {
      real: "",
      short:"",
      messageNotif:false,
      message:"",

      items:[],

      dataModal:[],
      copied:false,

      currentPage:1,
      perPage:5
    };
  },
  methods:{
    postUrl(){
      fetch("https://tautan.loophole.site/url", { 
        method: "POST",
        headers:{
          // "Access-Control-Allow-Origin":"*",
          "Content-Type":"application/json",
          // "Access-Control-Allow-Methods":"OPTIONS, GET, POST, PUT"
        },
        body: JSON.stringify({
            realUrl:this.real,
            shortUrl:this.short
        })
      })
      .then(async response => {
        const data = await response.json()
        console.log(data)
        this.items.push({
          "realUrl": this.real,
          "shortUrl": this.short
        })

        this.real=""
        this.short=""
        this.message=data
        this.messageNotif=true
        setTimeout(() => {
          this.messageNotif = false;
        }, 2000);


      })
      .catch( err=> {
        console.error(err)
      })
    },
    getUrl(){
      fetch(`https://tautan.loophole.site/url`,{
        headers:{
          "Content-Type":"application/json"
        }
      })
      .then(async response=>{
        const urls = await response.json()
        // console.log(urls)
        await urls.map(u=>{
          this.items.push(u)
        })
        console.items
      })
    },
    copyURL(mytext) {
      navigator.clipboard.writeText(mytext)
      .then(() => {
        this.copied = true;
        setTimeout(() => {
          this.copied = false;
        }, 2000);
      })
      .catch(err => {
        console.error('Failed to copy text:', err);
      });
    },
    detailUrl(url){
      fetch("https://tautan.loophole.site/detail/"+url,{
        headers:{
          "Content-Type":"application/json"
        }
      })
      .then(async response=>{
        const arr = await response.json()
        console.log(arr)
        arr.map(a=>{
          this.dataModal.push(a)
        })
      })
      .catch(err => {
        console.error('Failed to copy text:', err);
      });

    },
    deleteModal(){
      this.dataModal = []
    }
  },
  computed: {
    rows() {
      return this.items.length
    },
    maxPage() {
      let res = Math.ceil(this.rows/this.perPage)
      return res
    },
    itemToShow(){
      const startIndex = (this.currentPage - 1) * this.perPage;
      const endIndex = startIndex + this.perPage;

      return this.items.slice(startIndex, endIndex)
    },
    tautanAsli(){
      return this.dataModal[0].realUrl
    }
  },
  created(){
    this.getUrl()
  }
};


</script>

<template>
  
  <header class="mt-5">
    <h1>Tautan</h1>
  </header>

  <main >
    <textarea name="Text1" cols="100" rows="5" placeholder="enter url here" v-model="real"></textarea>
    <!-- <input class="m-2" type="textarea" name="realUrl" id="realUrl" placeholder="starts with www...." required> -->
    <label class="m-2" for="">to</label>
    <div class="input-group">
      <span class="input-group-text" id="basic-addon3">tautan.loophole.site/</span>
      <input v-model="short" type="text" class="form-control border-black" aria-describedby="basic-addon3 basic-addon4" name="shortUrl" id="shortUrl" placeholder="short version" required>    
      <span>
        <button class="btn btn-success m-3 form-control" type="button" @click="postUrl">Shorten</button>
      </span>
    </div>
    
    <div class="border rounded-1 bg-success-subtle" v-if="messageNotif">
      <h4 class="p-2">
        {{ message }}
      </h4>
    </div>

    <div class="overflow-auto my-5" id="pagination">
      <div class="pagination">        
        <li class="page-item"><button class="page-link" @click="currentPage--" :disabled="currentPage==1" >Previous</button></li>
        <li class="page-item"><a class="page-link" >{{currentPage}}</a></li>
        <li class="page-item"><button class="page-link" @click="currentPage++" :disabled="currentPage==maxPage" >Next</button></li>      </div>
      <div class="pagination">
        <li class="page-item"><a class="page-link" >item to show</a></li>        
        <li class="page-item"><input type="number" class="page-link" v-model="perPage"></input></li>
      </div>
    </div>
    <transition name="fade" class="mt-3">
        <div v-if="copied" class="copied-message">Copied!</div>
    </transition>
    
    <div class="url mb-5">
      <div class="border rounded-1 border-black my-3 p-2" v-for="d in itemToShow" id="urlItems" >
        <p >
          tautan.loophole.site/{{ d['shortUrl'] }} 
          <a class="fa fa-icon" @click="copyURL(`tautan.loophole.site/${d['shortUrl']}`)">Copy</a> 
        </p>
        <button class="btn btn-warning" data-bs-toggle="modal" data-bs-target="#detailModal" @click="detailUrl(`${d['shortUrl']}`)">detail</button>
      </div>
    </div>
  </main> 

  <!-- Modal -->
  <div class="modal fade" id="detailModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h1 class="modal-title fs-5" id="exampleModalLabel">detail</h1>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="deleteModal"></button>
        </div>
        <div class="modal-body">
          <p>tautan asli : {{ tautanAsli }}</p>
          <table class="table">
            <thead>
              <tr>
                <td>Tanggal</td>
                <td>akses berapa kali</td>
              </tr>
            </thead>
            <tbody>
              <tr  v-for="data in dataModal">
                  <td>
                    {{data['date'].slice(0,10)}}
                  </td>
                  <td>
                    {{data['count']}}
                  </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" @click="deleteModal">Close</button>
        </div>
      </div>
    </div>
  </div>


</template>

<style scoped>
header {
  display: grid;
  justify-content: center;
}
main{
  display: grid;
  align-items: center;
  justify-content: center;
}

#realUrl{
  height: 300px;
  width: 1000px;
}

#urlItems{
  display: flex;
  justify-content: space-between;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}
.copied-message {
  position: sticky; 
  top: 0;
  padding: 5px 16px;
  background-color: #4CAF50;
  color: white;
  border-radius: 4px;
  font-size: 14px;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}
#pagination{
  display: flex;
  justify-content: space-between;
}
</style>
