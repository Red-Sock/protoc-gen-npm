/* eslint-disable */
// @ts-nocheck

/**
 * This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
 */

import * as ExampleApiCommonP from "./common_p.pb";
import * as fm from "./fetch.pb";


export class ControlPlane {
  static DoNothing(this:void, req: ExampleApiCommonP.Empty, initReq?: fm.InitReq): Promise<ExampleApiCommonP.Empty> {
    return fm.fetchRequest<ExampleApiCommonP.Empty>(`/example_api.ControlPlane/DoNothing`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
}