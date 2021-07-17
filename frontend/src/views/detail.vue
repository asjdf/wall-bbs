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
          <el-card>
            <el-page-header @back="goBack" content="详情页面"/>
          </el-card>
          <div class="posts-list">
            <el-card class="posts-list-card">
              <!--              <el-dropdown class="management" v-if="right==1">-->
              <!--                <span class="el-dropdown-link">-->
              <!--                  管理<i class="el-icon-arrow-down el-icon&#45;&#45;right"></i>-->
              <!--                </span>-->
              <!--                <el-dropdown-menu slot="dropdown">-->
              <!--                  <el-dropdown-item><p @click="deletePost(blog.id)">删除</p></el-dropdown-item>-->
              <!--                </el-dropdown-menu>-->
              <!--              </el-dropdown>-->

              <!--              <div v-html="compiledMarkdown(blog.content)" class="post-content"/>-->
              <div v-html="parentInfo.content" class="post-content"/>
              <!--              <el-skeleton />-->
              <el-divider class="divider" content-position="right">{{ parentInfo.user_name }}</el-divider>
              <p class="post-time">{{ parentInfo.post_time }}</p>
            </el-card>

            <el-card
                class="posts-list-card"
            >
              <quill-editor
                  class="editor"
                  v-model="content"
                  :options="editorOption"
              />
              <el-row type="flex" justify="space-between" align="middle">
                <el-checkbox v-model="anonymous">匿名  </el-checkbox>
                <el-button type="primary" size="medium" round @click="postContent" :loading="isPosting">发送</el-button>
              </el-row>
            </el-card>


            <el-card
                v-for="comment in comments"
                :key="comment.content"
                :timestamp="comment.ptime"
                class="posts-list-card"
            >
              <!--              <el-dropdown class="management" v-if="right==1">-->
              <!--                <span class="el-dropdown-link">-->
              <!--                  管理<i class="el-icon-arrow-down el-icon&#45;&#45;right"></i>-->
              <!--                </span>-->
              <!--                <el-dropdown-menu slot="dropdown">-->
              <!--                  <el-dropdown-item><p @click="deletePost(blog.id)">删除</p></el-dropdown-item>-->
              <!--                </el-dropdown-menu>-->
              <!--              </el-dropdown>-->

              <!--              <div v-html="compiledMarkdown(blog.content)" class="post-content"/>-->
              <div v-html="comment.content" class="post-content" @click="goDetail(comment.pid)"/>
              <!--              <el-skeleton />-->
              <el-divider class="divider" content-position="right">{{ comment.user_name }}</el-divider>
              <p class="post-time">{{ comment.post_time }}</p>
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
export default {
  name: "detail",
  data() {
    return {
      anonymous: false,
      isPosting: false,
      totalPages : 1,
      currentPage: 1,
      parentInfo: {},
      comments: [],
      content: '',
      editorOption: {
        theme: 'snow',
        modules: {
          toolbar: {
            container: [
              ['bold','link', 'image']
            ],
            handlers: {
              // 'image': function () {
              //   QuillWatch.emit(this.quill.id)
              // }
            }
          },
          // ImageExtend: {
          //   loading: true,
          //   name: 'img',
          //   action: 'https://github.surmon.me/images/',
          //   response: (res) => {
          //     return testImageUrl
          //   }
          // }
        }
      }
    }
  },
  methods: {
    goBack() {
      window.history.length > 1 ? this.$router.go(-1) : this.$router.push('/')
    },
    getPostList() {
      this.$ajax.post('/posts/list', {
        parent_id: String(this.$route.params.pid),
        page_num: this.currentPage
      }).then((response) => {
        this.parentInfo = response.data.data.parent_post
        this.comments = response.data.data.child_list;
      })
    },
    goDetail(pid) {
      this.$router.push({ name: 'detail', params: { pid: pid } })
    },
    backTop () {
      window.scrollTo(0,0);
    },
    onPageChange(){
      this.getPostList();
      this.backTop();
    },
    postContent() {
      this.isPosting = true
      this.$ajax.post('/posts/new', {
        parent_id: String(this.$route.params.pid),
        content: this.content,
        anonymous: this.anonymous,
      }).then((response) => {
        if(response.data.code == 20000){
          this.$message({
            message: '发布成功',
            type: 'success'
          });
        }else{
          this.$message({
            message: response.data.msg,
            type: 'error'
          });
        }
        // eslint-disable-next-line no-unused-vars
      }).catch(err => {
        this.$message({
          message: "发布遇到其他错误",
          type: 'error'
        });
      })
      this.isPosting = false
    }
  },
  created() {
    this.getPostList();
  }
}
</script>

<style scoped>
  .el-row {
    width: 100%;
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