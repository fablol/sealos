import type { NextApiRequest, NextApiResponse } from 'next';
import { BadRequestResp, NotFoundResp, JsonResp } from '../response';
import { CheckIsInCluster, K8sApi, ListPods, ReplaceInCLuster } from '../../../lib/kubernetes';

export default async function handler(req: NextApiRequest, resp: NextApiResponse) {
  if (req.method !== 'GET') {
    return BadRequestResp(resp);
  }

  const { kubeconfig } = req.body;
  if (kubeconfig === '') {
    return NotFoundResp(resp);
  }

  const kubeconfigR = ReplaceInCLuster(kubeconfig);
  const kc = K8sApi(kubeconfigR);

  const res = await ListPods(kc, 'sealos');
  JsonResp(res.body, resp);
}
