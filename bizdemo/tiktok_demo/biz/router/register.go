/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	comment "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/router/comment"
	favorite "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/router/favorite"
	feed "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/router/feed"
	message "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/router/message"
	publish "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/router/publish"
	relation "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/router/relation"
	user "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/router/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	message.Register(r)

	relation.Register(r)

	comment.Register(r)

	favorite.Register(r)

	feed.Register(r)

	publish.Register(r)

	user.Register(r)

}