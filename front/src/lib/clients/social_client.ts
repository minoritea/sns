import {createConnectTransport, createPromiseClient} from '@bufbuild/connect-web';
import {SocialService} from '$lib/proto/social_connectweb';

export default createPromiseClient(SocialService, createConnectTransport({baseUrl: '/rpc'}));
