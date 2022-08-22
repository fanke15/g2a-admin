APP_NAME=g2aAdmin
if [ -n $1 ]; then
  APP_NAME=$1
fi

CUR_DIR=$(pwd)
RUN_PATH=app
RUN_DIR=${CUR_DIR}/${RUN_PATH}
TARGET_OUT_DIR="deploy"
TARGET_OUT_PATH=${CUR_DIR}/$TARGET_OUT_DIR
TARGET_OUT_NAME=${TARGET_OUT_PATH}/${APP_NAME}

buildCmd() {
  if [ ! -d "${TARGET_OUT_DIR}" ]; then
    mkdir -p ${TARGET_OUT_DIR}
  fi

  echo "从[${RUN_DIR}]开始构建${APP_NAME}..."
  $(go build -o ${TARGET_OUT_NAME} ${RUN_DIR})
}

buildCmd