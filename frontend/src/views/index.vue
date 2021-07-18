<template>

  <el-container>
<!--    <el-row>-->
<!--      <el-carousel :interval="5000" arrow="always">-->
<!--        <el-carousel-item v-for="item in 4" :key="item">-->
<!--          <h3>{{ item }}</h3>-->
<!--        </el-carousel-item>-->
<!--      </el-carousel>-->
<!--    </el-row>-->
    <el-row type="flex" justify="center">
      <el-col :sm="22" :md="20" :lg="18">
        <el-col :xs="24" :sm="16" :md="17">
          <div class="board-list">
            <el-row>
              <el-col
                  :xs="12" :sm="6" :md="6" :lg="4" :xl="1"
                  v-for="board in boards"
                  :key="board.board_name"
              >
                <div @click="goBoard(board.board_id)">
                  <el-card class="board-list-card">
                    <el-avatar :size="70" :src="board.bpic">{{board.board_name}}</el-avatar>
                    <el-divider class="divider">{{board.board_name}}</el-divider>
                  </el-card>
                </div>

              </el-col>


            </el-row>
          </div>

          <div class="posts-list">
            <el-card
                v-for="blog in blogs"
                :key="blog.content"
                :timestamp="blog.ptime"
                class="posts-list-card"
            >
<!--              v-if="this.$store.state.right===1||blog.uid===this.$store.state.uid"-->
              <el-dropdown class="management" v-if="displayManagement(blog.uid)">
                <span class="el-dropdown-link">
                  管理<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item><p @click="deletePost(blog.pid)">删除</p></el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>

<!--              <div v-html="compiledMarkdown(blog.content)" class="post-content"/>-->
              <div v-html="blog.content" class="post-content" @click="goDetail(blog.pid)"/>
<!--              <el-skeleton />-->
              <el-divider class="divider" content-position="right">{{blog.user_name}}</el-divider>
              <p class="post-time">{{ blog.post_time }}</p>
            </el-card>
          </div>

          <el-pagination
              v-model="currentPage"
              layout="prev, pager, next"
              :page-count="totalPages"
              :current-page.sync="currentPage"
              :hide-on-single-page="true"
              @current-change="onPageChange"
              background
          />
        </el-col>
        <el-col :sm="8" :md="7" class="hidden-xs-only right-list">
          <el-card>
            <p>施工中，敬请期待</p>
          </el-card>
        </el-col>
      </el-col>
    </el-row>
  </el-container>
</template>

<script>
// import HelloWorld from "@/components/HelloWorld";
export default {
  name: "index",
  data() {
    return {
      totalPages : 1,
      currentPage: 1,
      blogs: [],
      boards: []
    }
  },
  components: {
    // HelloWorld
  },
  methods: {
    getPostList() {
      this.$ajax.post('/posts/list', {
        parent_id: "",
        page_num: this.currentPage,
        desc: true
      }).then((response) => {
        this.totalPages = response.data.data.total_page;
        this.blogs = response.data.data.child_list;
      })
    },
    getBoardList() {
      this.$ajax.get('/boards/list').then((response) => {
        this.boards = response.data.data.list;
      })
    },
    goDetail(pid) {
      this.$router.push({ name: 'detail', params: { pid: pid } })
    },
    goBoard(bid) {
      this.$router.push({ name: 'board', params: { bid: bid } })
    },
    backTop () {
      window.scrollTo(0,0);
    },
    onPageChange(){
      this.getPostList();
      this.backTop();
    },
    displayManagement(uid) {
      return this.$store.state.right==1||uid==this.$store.state.uid
    },
    deletePost(pid) {
      this.$ajax.post('/posts/delete', {
        pid: pid,
      }).then((response) => {
        if(response.data.code === 20000){
          this.$message({
            message: '删除成功',
            type: 'success'
          });
          this.getPostList();
        }else{
          this.$message({
            message: response.data.msg,
            type: 'error'
          });
        }
        // eslint-disable-next-line no-unused-vars
      }).catch(err => {
        this.$message({
          message: "删除遇到其他错误",
          type: 'error'
        });
      })
    }
  },
  created() {
    this.getPostList();
    this.getBoardList();
  }
}
</script>

<style scoped>
  .el-row {
    width: 100%;
  }
  .board-list-card {
    padding-bottom: 0px;
    margin: 0.5em;
  }
  .board-list-card .divider {
    margin: 1em 0 0 0;
  }
  .posts-list-card {
    margin-bottom: 1em;
  }
  .posts-list-card p {
    text-align: left;
    margin: 1em 0em;
  }
  .posts-list-card .divider {
    margin: 1em 0;
  }
  .post-content {
    padding: 20px 0 0;
  }
  .post-content >>> img {
    width: 93%;
  }
  .posts-list-card .management {
    float: right;
  }
  .posts-list-card .post-time{
    text-align: right;
    color: #2c3e50;
    font-size: 50%;
    margin: 0;
  }
  .right-list {
    padding-left: 2em;
  }
</style>