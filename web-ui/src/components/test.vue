<!--<div  class="span-left" style="margin-left: 10%">
    <p style="color: #52c41a!important;font-size: 15pt;">/test/abc</p>
    <p>
        <Tag color="geekblue">SpringMvc</Tag>&nbsp;
        <Tag color="green">2ms</Tag>
    </p>
</div>-->

<script>
    export default {
        name: "test",
        props : {
          spanDeepth: {
              type:Number ,
              default: 0
          },
          isRoot: {
              type: Boolean ,
              default: false ,
          }
        },
        render(h) {
            return h('div',{},this.renderSpan(h,this.spanDeepth,this.isRoot))
        },
        methods:{
            renderSpan(h ,deepth,isRoot) {
                let p1 = h('p',{
                    style:{
                        color: '#52c41a!important',
                        fontSize: "15pt"
                    }
                },'/test/abc') ;

                let p2 = h('p',{},[
                    h('Tag',{props:{
                            color: "geekblue"
                        }},'SpringMvc'),
                    h('Tag',{props:{
                            color: "green"
                        }},'2ms'),
                ])
                if(deepth == 1) {
                    console.log("1231")
                }
                let spanClass = 'span-left' ;
                if(deepth == this.spanDeepth && !isRoot) {
                    spanClass = 'span-left-root'
                }else if (deepth > 0) {
                    spanClass = 'span-left-border'
                }
                let pDiv =  h('div',{
                    class: spanClass,
                }, deepth == 0 ? [p1,p2] : this.renderSpan(h,deepth-1))  ;
                return [pDiv] ;
            }
        }
    }
</script>

<style scoped>
    .span-left {
        border-left: 2px solid gray ;
        padding-bottom: 5px;
        padding-left: 10px ;
    }
    .span-left-border {
        border-left: 2px solid gray ;
        padding-left: 20px ;
    }
    .span-left-root {
        padding-left: 20px ;
    }
</style>