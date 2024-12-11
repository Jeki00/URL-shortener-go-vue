<script>
export default {
    data(){
        return{
            items:[],
            dataModal:[],
            copied:false,
            currentPage:1,
            perPage:5,
            tautan:'waiting'
        }
    },
    methods:{
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
                await arr.map(a=>{
                    this.dataModal.push(a)
                })
                this.tautan = this.dataModal[0].realUrl
                // console.log(this.dataModal[0])

            })
            .catch(err => {
                console.error('Failed to copy text:', err);
            });
        },
        deleteModal(){
            this.dataModal = []
            this.tautan="waiting"
        }
    },
    computed:{
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
    },
    created(){
        this.getUrl()
    }
}
</script>

<template>
    <transition name="fade" class="mt-3">
        <div v-if="copied" class="copied-message">Copied!</div>
    </transition>
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
        
    <div class="url mb-5 border p-5 shadow-sm">
        <div class="border rounded-1 border-black my-3 p-2" v-for="d in itemToShow" id="urlItems" >
            <p >
            tautan.loophole.site/{{ d['shortUrl'] }} 
            <a class="fa fa-icon" @click="copyURL(`tautan.loophole.site/${d['shortUrl']}`)">Copy</a> 
            </p>
            <button class="btn btn-warning" data-bs-toggle="modal" data-bs-target="#detailModal" @click="detailUrl(`${d['shortUrl']}`)">detail</button>
        </div>
    </div>   

    <!-- Modal -->
    <div class="modal fade" id="detailModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h1 class="modal-title fs-5" id="exampleModalLabel">Detail</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="deleteModal"></button>
                </div>
                <div class="modal-body">
                    <p>tautan asli : {{ tautan }}</p>
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

<style>
#pagination{
    display: flex;
    justify-content: space-between;
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
#urlItems{
    display: flex;
    justify-content: space-between;
}

</style>
