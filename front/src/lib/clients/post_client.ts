import { createConnectTransport, createPromiseClient } from '@bufbuild/connect-web';
import { PostService } from '$lib/proto/post_connectweb';
import type { Post } from '$lib/proto/post_pb';

export default createPromiseClient(PostService, createConnectTransport({ baseUrl: '/rpc' }));
export type { Post };
