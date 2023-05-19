<template>
  <a-row>
    <a-col :xs="{ span: 0, offset: 0 }"  :lg="{ span: 16, offset: 0 }" class="x1">
    </a-col>
    <!-- <c-col :span="2"></c-col> -->
    <a-col style="background-color: rgb(255, 255, 255);" :xs="{ span: 24, offset: 0 }"  :lg="{ span: 6, offset: 0 }">
        <div>
          <a-form class="login-form" ref="formRef" :model="formState">
            <a-form-item label="账号" name="username">
              <a-input class="inputBox" v-model:value="formState.username"></a-input>
            </a-form-item>
            <a-form-item label="密码" name="username">
              <a-input-password class="inputBox" @keyup.enter="onSubmit" v-model:value="formState.password">
              </a-input-password>
            </a-form-item>

            <a-form-item>
              <a-button class="submit" type="primary" @click="onSubmit">登录</a-button>
            </a-form-item>
          </a-form>
        </div>
      </a-col>
  </a-row>
</template>

<script lang="ts">
import { defineComponent, onMounted, toRaw, ref, reactive } from "vue";
import type { UnwrapRef } from "vue";
import axios from "axios";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { message } from "ant-design-vue";

interface FormState {
  username: string;
  password: string;
}

export default defineComponent({
  components: {},
  setup() {
    const formRef = ref();
    const router = useRouter();
    const store = useStore();
    const formState: UnwrapRef<FormState> = reactive({
      username: "admin",
      password: "admin",
    });



    onMounted(() => {
      const token = localStorage.getItem('token') as string
      if (token?.length > 10) {
        router.push('/adm/knowledge/treelist');

      }
    });


    const onSubmit = () => {
      axios({
        url: "/api/login",
        method: "POST",
        data: toRaw(formState),
      }).then((res) => {
        if (res.data.code == 400) {
          message.error("用户名或密码错误??")
        } else if (res.data.code == 200) {
          const token = res.data.data.token;
          localStorage.setItem("token", token);
          router.push("/adm/dashboard");
        } else if (res.data.code == 500) {
          message.error(res.data.msg)
        }

      });
    };

    return {
      formRef,
      formState,
      onSubmit,
    };
  },
});
</script>

<style>
.login-form {
  margin: 300px 70px 0 70px;
  min-height: 540px;
}


.submit {
  width: 80%;
  height: 45px;
  margin: 0 0 0 55px;
  font-size: 16px;
}

/* 输入框 */
.inputBox {
  height: 45px;
}
.x1{
  background-repeat: no-repeat;
  /* background-color: blue; */
  background-image: url("../assets/bg.png");
  background-size: 100% 100%;
}

/* 输入框内左边距50px */
.ant-input-affix-wrapper .ant-input:not(:first-child) {
  padding-left: 50px;
}
</style>
