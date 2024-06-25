<template>
    <div class="common-layout">
        <el-container class="tutor-container">
            <el-header>
                <el-icon>
                    <UserFilled />
                </el-icon>
                <span>
                    家教信息管理(学生端)
                </span>
                <el-button type="info" class="ce" v-on:click="out">退出</el-button>
            </el-header>
            <el-container>
                <el-aside width="200px">
                    <el-menu router="true" class="el-menu-vertical-demo" active-text-color="#071a52"
                        background-color="#dbe2ef">

                        <el-menu-item index="/students/profile">
                            <el-icon>
                                <Grid />
                            </el-icon>
                            <span>学生信息发布</span>
                        </el-menu-item>

                        <el-menu-item index="/matching_student">
                            <el-icon>
                                <Grid />
                            </el-icon>
                            <span>家教-学生推荐匹配</span>
                        </el-menu-item>
                        <el-menu-item index="/notifications_student">
                            <el-icon>
                                <Grid />
                            </el-icon>
                            <span>消息通知</span>
                        </el-menu-item>

                    </el-menu>
                </el-aside>
                <el-main>

                    <router-view></router-view>

                </el-main>
            </el-container>
        </el-container>
    </div>
</template>

<script>
// @ is an alias to /src

export default {
    methods: {
        out() {
            this.$router.push('/')
        }
    }
}
</script>
<style lang="less" scoped>
.common-layout {
    height: 100%;

}

.tutor-container {
    height: 100%;
}

.el-icon {
    padding: 10px;
    display: flex;
    align-items: center;
    height: 70%;

}

.ce {
    display: inline-block;
    width: 50px;
    height: 30px;
    background: #c00;

}

.el-header {
    background-color: #112d4e;
    display: flex;
    justify-content: space-between;
    font-size: 20px;
    color: aliceblue;

    >span {
        display: flex;
        align-items: center;
    }


}

.el-aside {
    background-color: #005691;
}

.el-main {
    background-color: #f9f7f7;
}
</style>