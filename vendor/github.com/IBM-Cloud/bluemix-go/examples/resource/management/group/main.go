package main

import (
	"flag"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var resourcegrp string
	flag.StringVar(&resourcegrp, "name", "", "Name of the group")

<<<<<<< HEAD
	var resourcequota string
	flag.StringVar(&resourcequota, "quota", "", "Name of the group")

	flag.Parse()

	if resourcegrp == "" || resourcequota == "" {
=======
	flag.Parse()

	if resourcegrp == "" {
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	client, err := management.New(sess)
	if err != nil {
		log.Fatal(err)
	}

<<<<<<< HEAD
	resQuotaAPI := client.ResourceQuota()

	quota, err := resQuotaAPI.FindByName(resourcequota)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Quota Defination Details :", quota)

=======
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	resGrpAPI := client.ResourceGroup()

	resourceGroupQuery := management.ResourceGroupQuery{
		Default: true,
	}

	grpList, err := resGrpAPI.List(&resourceGroupQuery)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resource group default Details :", grpList)

	var grpInfo = models.ResourceGroup{
<<<<<<< HEAD
		Name:    resourcegrp,
		QuotaID: quota[0].ID,
=======
		Name: resourcegrp,
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	}

	grp, err := resGrpAPI.Create(grpInfo)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resource group create :", grp)

<<<<<<< HEAD
	grps, err := resGrpAPI.FindByName(nil, grp.Name)
=======
	grps, err := resGrpAPI.FindByName(nil, resourcegrp)
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resource group Details :", grps[0])

	grp, err = resGrpAPI.Get(grp.ID)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resource group Details by ID:", grp)

	var updateGrpInfo = management.ResourceGroupUpdateRequest{
<<<<<<< HEAD
		Name:    "default",
		QuotaID: quota[0].ID,
=======
		Name: "default",
>>>>>>> 39014884d69db9425c92363e89383b38bba01fbe
	}

	grp, err = resGrpAPI.Update(grp.ID, &updateGrpInfo)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Resource group update :", grp)

	err = resGrpAPI.Delete(grp.ID)
	if err != nil {
		log.Fatal(err)
	}

}
