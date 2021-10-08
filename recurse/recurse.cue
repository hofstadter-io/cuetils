package recurse

import "list"

#RecurseN: {
	#maxiter: uint | *20
	#funcFactory: {
		#next: _
		#func: _
	}

	for k, v in list.Range(0, #maxiter, 1) {
		#funcs: "\(k)": (#funcFactory & {#next: #funcs["\(k+1)"]}).#func
	}
	#funcs: "\(#maxiter)": null

	#funcs["0"]
}
